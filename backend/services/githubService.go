package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type GithubService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GitHubResponseToken, error)
	GetUserInfo(accessToken string) (schemas.GithubUserInfo, error)
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64)
	FindReactionByName(name string) func(workflowId uint64, accessToken []schemas.ServiceToken)
}

type githubService struct {
	repository repository.GithubRepository
	userService UserService
	reactionResponseDataService ReactionResponseDataService
}

func NewGithubService(
	repository repository.GithubRepository,
	userService UserService,
	reactionResponseDataService ReactionResponseDataService,
	) GithubService {
	return &githubService{
		repository: repository,
		userService: userService,
		reactionResponseDataService: reactionResponseDataService,
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
	var resultToken schemas.GitHubResponseToken
	err = json.NewDecoder(response.Body).Decode(&resultToken)
	if err != nil {
		return schemas.GitHubResponseToken{}, err
	}
	response.Body.Close()
	return resultToken, nil
}

func (service *githubService) GetUserInfo(accessToken string) (schemas.GithubUserInfo, error) {
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	request.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	result := schemas.GithubUserInfo{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	response.Body.Close()
	return result, nil
}

func (service *githubService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64) {
	switch name {
	case string(schemas.GithubPullRequest):
		return service.LookAtPullRequest
	default:
		return nil
	}
}

func (service *githubService) FindReactionByName(name string) func(workflowId uint64, accessToken []schemas.ServiceToken) {
	switch name {
	case string(schemas.GithubReactionListComments):
		return service.ListAllReviewComments
	default:
		return nil
	}
}

func (service *githubService) LookAtPullRequest(channel chan string, option string, workflowId uint64) {
	time.Sleep(30 * time.Second)
	fmt.Printf("LookAtPullRequest\n")
	channel <- "LookAtPullRequest"
}

func (service *githubService) ListAllReviewComments(workflowId uint64, accessToken []schemas.ServiceToken) {
	request, err := http.NewRequest("GET", "https://api.github.com/repos/Epitouche/Area51/pulls/comments", nil)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}
	defer response.Body.Close()

	result := []schemas.GithubListCommentsResponse{}
	savedResult := schemas.ReactionResponseData{
		WorkflowId: workflowId,
	}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}

	jsonValue, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	savedResult.ApiResponse = json.RawMessage(jsonValue)

	if err != nil {
		fmt.Println(err)
	}
	service.reactionResponseDataService.Save(savedResult)
}