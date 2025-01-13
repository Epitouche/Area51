package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area51/toolbox"
)

type MicrosoftController interface {
	RedirectionToMicrosoftService(ctx *gin.Context, path string) (string, error)
	ServiceMicrosoftCallback(ctx *gin.Context, path string) (string, error)
}

type microsoftController struct {
}

func NewMicrosoftController() MicrosoftController {
	return &microsoftController{}
}

func (controller *microsoftController) RedirectionToMicrosoftService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("MICROSOFT_CLIENT_ID")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	state, err := toolbox.GenerateCSRFToken()
	if err != nil {
		return "", err
	}
	redirectUri := fmt.Sprintf("%s%s/callback", appAdressHost, appPort)
	authUrl := fmt.Sprintf("https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=%s&response_type=code&scope=openid&redirect_uri=%s&state=%s", clientId, redirectUri, state)
	return authUrl, nil
}

func (controller *microsoftController) ServiceMicrosoftCallback(ctx *gin.Context, path string) (string, error) {
	return "microsoft_token", nil
}