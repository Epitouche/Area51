package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
)

type MicrosoftApi struct {
	controller controllers.MicrosoftController
}

func NewMicrosoftApi(controller controllers.MicrosoftController) *MicrosoftApi {
	return &MicrosoftApi{
		controller: controller,
	}
}

func (api *MicrosoftApi) RedirectToMicrosoft(ctx *gin.Context, path string) {
	if authURL, err := api.controller.RedirectionToMicrosoftService(ctx, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, schemas.OAuthConnectionResponse{
			ServiceAuthenticationUrl: authURL,
		})
	}
}

func (api *MicrosoftApi) HandleMicrosoftTokenCallback(ctx *gin.Context, path string) {
	if microsoft_token, err := api.controller.ServiceMicrosoftCallback(ctx, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": microsoft_token})
	}
}
