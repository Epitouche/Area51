package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-github/v67/github"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type GithubService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GitHubResponseToken, error)
	GetUserInfo(accessToken string) (schemas.GithubUserInfo, error)
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64)
	FindReactionByName(name string) func(workflowId uint64)
}

type githubService struct {
	repository repository.GithubRepository
}

func NewGithubService(repository repository.GithubRepository) GithubService {
	return &githubService{
		repository: repository,
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

func (service *githubService) FindReactionByName(name string) func(workflowId uint64) {
	switch name {
	case string(schemas.GithubReactionCreateNewRelease):
		return service.CreateNewRelease
	default:
		return nil
	}
}

var nbPR int

func (service *githubService) LookAtPullRequest(channel chan string, option string, workflowId uint64) {
	ctx := context.Background()
	client := github.NewClient(nil)
	// var options schemas.GithubPullRequestOptions
	// err := json.NewDecoder(strings.NewReader(option)).Decode(&options)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	pullRequests, _, err := client.PullRequests.List(ctx, "Karumapathetic", "https://github.com/Karumapathetic/testAREA", nil)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}
	if nbPR < len(pullRequests) {
		fmt.Println("Trigger reaction")
		nbPR = len(pullRequests)
	}
	time.Sleep(30 * time.Second)
}

func (service *githubService) CreateNewRelease(workflowId uint64) {
	time.Sleep(30 * time.Second)
	fmt.Printf("CreateNewRelease\n")
}