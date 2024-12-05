package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/toolbox"
)

type GithubController interface {
	RedirectionToGithubService(ctx *gin.Context, path string) (string, error)
}

type githubController struct {
	service services.GithubService
	
}

func NewGithubController()


func (controller *githubController) RedirectionToGithubService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("GITHUB_CLIENT_ID")
	appPort := toolbox.GetInEnv("APP_PORT")
	appAdressHost := toolbox.GetInEnv("APP_ADRESS_HOST")

	state, err := toolbox.GenerateCSRFToken()
	if err != nil {
		return "", err
	}

	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	redirectUri := appAdressHost + appPort + path
	authUrl := "https://github.com/login/oauth/authorize" +
		"?client_id=" + clientId +
		"&response_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectUri +
		"&state=" + state
	return authUrl, nil
}
