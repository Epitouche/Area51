package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
)

type GoogleApi struct {
	controller controllers.GoogleController
}

func NewGoogleApi(controller controllers.GoogleController) *GoogleApi {
	return &GoogleApi{
		controller: controller,
	}
}

func (api *GoogleApi) RedirectToGoogle(ctx *gin.Context, path string) {
	if authURL, err := api.controller.RedirectionToGoogleService(ctx, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, schemas.OAuthConnectionResponse{
			ServiceAuthenticationUrl: authURL,
		})
	}
}

func (api *GoogleApi) HandleGoogleTokenCallback(ctx *gin.Context, path string) {
	if google_token, err := api.controller.ServiceGoogleCallback(ctx, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": google_token})
	}
}
