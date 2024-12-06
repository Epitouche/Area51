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
	fmt.Printf("response body: %v\n", response.Body)
	var resultToken schemas.GitHubResponseToken
	err = json.NewDecoder(response.Body).Decode(&resultToken)
	if err != nil {
		return schemas.GitHubResponseToken{}, err
	}
	response.Body.Close()
	fmt.Printf("resultToken: %v\n", resultToken)
	return resultToken, nil
}

func (service *githubService) GetUserInfo(accessToken string) (schemas.GithubUserInfo, error) {
	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	fmt.Printf("accessToken: %s\n", accessToken)
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
