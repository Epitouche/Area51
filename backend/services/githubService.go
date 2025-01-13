package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/go-github/v67/github"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type GithubService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GitHubResponseToken, error)
	// GetUserInfo(accessToken string) (schemas.GithubUserInfo, error)
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
	GetUserInfosByToken(accessToken string) func(*schemas.ServicesUserInfos)
}

type githubService struct {
	githubRepository            repository.GithubRepository
	tokenRepository             repository.TokenRepository
	userService                 UserService
	workflowRepository          repository.WorkflowRepository
	reactionRepository          repository.ReactionRepository
	reactionResponseDataService ReactionResponseDataService
	mutex                       sync.Mutex
}

func NewGithubService(
	githubRepository repository.GithubRepository,
	tokenRepository repository.TokenRepository,
	workflowRepository repository.WorkflowRepository,
	reactionRepository repository.ReactionRepository,
	reactionResponseDataService ReactionResponseDataService,
	userService UserService,
) GithubService {
	return &githubService{
		githubRepository:            githubRepository,
		tokenRepository:             tokenRepository,
		workflowRepository:          workflowRepository,
		reactionRepository:          reactionRepository,
		reactionResponseDataService: reactionResponseDataService,
		userService:                 userService,
	}
}

func (service *githubService) AuthGetServiceAccessToken(code string, path string) (schemas.GitHubResponseToken, error) {
	clientId := toolbox.GetInEnv("GITHUB_CLIENT_ID")
	clientSecret := toolbox.GetInEnv("GITHUB_SECRET")
	appPort := toolbox.GetInEnv("APP_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	redirectUri := appAdressHost + appPort + path
	apiUrl := "https://github.com/login/oauth/access_token"

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)

	req, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return schemas.GitHubResponseToken{}, err
	}

	req.URL.RawQuery = data.Encode()
	req.Header.Set("Accept", "application/json")
	client := &http.Client{
		Timeout: time.Second * 45,
	}

	response, err := client.Do(req)
	if err != nil {
		return schemas.GitHubResponseToken{}, err
	}

	resultToken := schemas.GitHubResponseToken{}

	err = json.NewDecoder(response.Body).Decode(&resultToken)
	if err != nil {
		return schemas.GitHubResponseToken{}, err
	}

	response.Body.Close()
	return resultToken, nil
}

// func (service *githubService) GetUserInfo(accessToken string) (schemas.GithubUserInfo, error) {
// 	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
// 	if err != nil {
// 		return schemas.GithubUserInfo{}, err
// 	}

// 	request.Header.Set("Authorization", "Bearer "+accessToken)
// 	client := &http.Client{}

// 	response, err := client.Do(request)
// 	if err != nil {
// 		return schemas.GithubUserInfo{}, err
// 	}

// 	result := schemas.GithubUserInfo{}

// 	err = json.NewDecoder(response.Body).Decode(&result)
// 	if err != nil {
// 		return schemas.GithubUserInfo{}, err
// 	}

// 	response.Body.Close()
// 	return result, nil
// }

func (service *githubService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	switch name {
	case string(schemas.GithubPullRequest):
		return service.LookAtPullRequest
	default:
		return nil
	}
}

func (service *githubService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	switch name {
	case string(schemas.GithubReactionListComments):
		return service.ListAllReviewComments
	default:
		return nil
	}
}

// var nbPR int

type transportWithToken struct {
	token string
}

func (t *transportWithToken) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	return http.DefaultTransport.RoundTrip(req)
}

func (service *githubService) LookAtPullRequest(channel chan string, option string, workflowId uint64, actionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()
	ctx := context.Background()

	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := service.userService.GetUserById(workflow.UserId)
	token := service.tokenRepository.FindByUserId(user)

	client := github.NewClient(&http.Client{
		Transport: &transportWithToken{token: token[len(token)-1].Token},
	})

	var actionData schemas.GithubPullRequestOptions
	err = json.Unmarshal([]byte(actionOption), &actionData)
	if err != nil {
		fmt.Println("Error parsing actionOption:", err)
		return
	}

	pullRequests, _, err := client.PullRequests.List(ctx, actionData.Owner, actionData.Repo, nil)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	existingRecords := service.githubRepository.FindByOwnerAndRepo(actionData.Owner, actionData.Repo)
	if existingRecords.UserId == 0 {
		service.githubRepository.Save(schemas.GithubPullRequestOptionsTable{
			UserId: user.Id,
			Repo:   actionData.Repo,
			Owner:  actionData.Owner,
			NumPR:  0,
		})
	}

	if existingRecords.NumPR != len(pullRequests) {
		actualRecords := service.githubRepository.FindByOwnerAndRepo(actionData.Owner, actionData.Repo)
		actualRecords.NumPR = len(pullRequests)
		workflow.ReactionTrigger = true
		service.workflowRepository.Update(workflow)
		service.githubRepository.UpdateNumPRs(actualRecords)
	}
	channel <- "Action workflow done"
}

func (service *githubService) ListAllReviewComments(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	// Iterate over all tokens
	for _, token := range accessToken {
		actualUser := service.userService.GetUserById(token.UserId)
		if token.UserId == actualUser.Id {
			actualWorkflow := service.workflowRepository.FindByUserId(actualUser.Id)
			for _, workflow := range actualWorkflow {
				if workflow.Id == workflowId {
					if !workflow.ReactionTrigger {
						fmt.Println("Trigger is already false, skipping reaction.")
						return
					}
				}
			}
		}
	}

	var reactionData schemas.GithubListAllReviewCommentsOptions
	err := json.Unmarshal([]byte(reactionOption), &reactionData)
	if err != nil {
		fmt.Println("Error parsing actionOption:", err)
		return
	}
	request, err := http.NewRequest("GET", "https://api.github.com/repos/"+reactionData.Owner+"/"+reactionData.Repo+"/pulls/comments", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Set("Accept", "application/vnd.github+json")
	for _, token := range accessToken {
		actualUser := service.userService.GetUserById(token.UserId)
		if token.UserId == actualUser.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer response.Body.Close()

	var result []schemas.GithubListCommentsResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	savedResult := schemas.ReactionResponseData{
		WorkflowId:  workflowId,
		ApiResponse: json.RawMessage{},
	}
	jsonValue, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error marshalling response:", err)
		return
	}
	savedResult.ApiResponse = jsonValue
	service.reactionResponseDataService.Save(savedResult)
	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
		return
	}
	workflow.ReactionTrigger = false
	service.workflowRepository.UpdateReactionTrigger(workflow)
	time.Sleep(1 * time.Minute)
}

func (service *githubService) GetUserInfosByToken(accessToken string) func(*schemas.ServicesUserInfos) {
	return func(userInfos *schemas.ServicesUserInfos) {
		request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
		if err != nil {
			return
		}

		request.Header.Set("Authorization", "Bearer "+accessToken)
		client := &http.Client{}

		response, err := client.Do(request)
		if err != nil {
			return
		}

		err = json.NewDecoder(response.Body).Decode(&userInfos.GithubUserInfos)
		if err != nil {
			return
		}

		response.Body.Close()
	}
}
