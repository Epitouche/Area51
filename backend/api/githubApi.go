package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
)


type GithubApi struct {
	controller controllers.GithubController
}

func NewGithubApi(controller controllers.GithubController) *GithubApi {
	return &GithubApi{
		controller: controller,
	}
}

func (api *GithubApi) RedirectToGithub(ctx *gin.Context, path string) {
	authURL, err := api.controller.RedirectionToGithubService(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"github_authentication_url": authURL})
	}
}

func (api *GithubApi) HandleGithubTokenCallback(ctx *gin.Context, path string) {
	github_token, err := api.controller.ServiceGithubCallback(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": github_token})
	}
}

func (api *GithubApi) StoreMobileToken(ctx *gin.Context) {
	token, err := api.controller.StoreMobileToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}