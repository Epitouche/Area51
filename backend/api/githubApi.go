package api

import (
	"area51/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GithubApi struct {
	githubTokenController controllers.GitHubController
}

func NewGithubAPI(githubTokenController controllers.GitHubController) *GithubApi {
	return &GithubApi{
		githubTokenController: githubTokenController,
	}
}

func (api *GithubApi) RedirectToGithub(ctx *gin.Context, path string) {
	authURL, err := api.githubTokenController.RedirectToGithub(ctx, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"github_authentication_url": authURL})
	}
}