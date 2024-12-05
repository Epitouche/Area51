package controllers

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"

	"area51/services"
	"area51/tools"
)

type GitHubController interface {
	RedirectToGithub(ctx *gin.Context, path string) (string, error)
}

type githubTokenControllerStruct struct {
	service services.GithubTokenService
	serviceUser services.UserService
}

func NewGithubTokenController(service services.GithubTokenService, serviceUser services.UserService) GitHubController {
	return &githubTokenControllerStruct{
		service: service,
		serviceUser: serviceUser,
	}
}

func (controller *githubTokenControllerStruct) RedirectToGithub(ctx *gin.Context, path string) (string, error) {
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	if clientID == "" {
		return "", errors.New("GITHUB_CLIENT_ID is not set")
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		return "", errors.New("APP_PORT is not set")
	}
	state, err := tools.GenerateCSRFToken()
	if err != nil {
		return "", errors.New("unable to generate CSRF token")
	}

	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	redirectURI := "http://localhost:" + appPort + path
	authURL := "https://github.com/login/oauth/authorize" +
		"?client_id=" + clientID +
		"&reponse_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectURI +
		"&state=" + state
	return authURL, nil
}
