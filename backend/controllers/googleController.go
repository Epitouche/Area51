package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

type GoogleController interface {
	RedirectionToGoogleService(ctx *gin.Context, path string) (string, error)
	ServiceGoogleCallback(ctx *gin.Context, path string) (string, error)
}

type googleController struct {
	service         services.GoogleService
	userService     services.UserService
	servicesService services.ServicesService
}

func NewGoogleController(
	service services.GoogleService,
	userService services.UserService,
	servicesService services.ServicesService,
) GoogleController {
	return &googleController{
		service:         service,
		userService:     userService,
		servicesService: servicesService,
	}
}

func (controller *googleController) RedirectionToGoogleService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("GOOGLE_CLIENT_ID")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	state, err := toolbox.GenerateCSRFToken()
	if err != nil {
		return "", err
	}
	redirectUri := fmt.Sprintf("%s%s/callback", appAdressHost, appPort)
	authUrl := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=openid&redirect_uri=%s&state=%s", clientId, redirectUri, state)
	return authUrl, nil
}

func (controller *googleController) ServiceGoogleCallback(ctx *gin.Context, path string) (string, error) {
	// var isAlreadyRegistered bool = false
	var codeCredentials schemas.OAuth2CodeCredentials
	err := json.NewDecoder(ctx.Request.Body).Decode(&codeCredentials)
	if err != nil {
		return "", err
	}
	if codeCredentials.Code == "" {
		return "", nil
	}
	if codeCredentials.State == "" {
		return "", nil
	}
	googleServiceToken, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
	fmt.Printf("googleServiceToken: %++v\n", googleServiceToken)
	if err != nil {
		return "", err
	}
	authHeader := ctx.GetHeader("Authorization")

	if authHeader != "" && len(authHeader) >= len("Bearer ") {
		token := authHeader[len("Bearer "):]
		user, err := controller.userService.GetUserInfos(token)
		if err != nil {
			return "", err
		}
		if user.Username != "" {
			err := controller.userService.AddServiceToUser(user, schemas.ServiceToken{
				Token:     googleServiceToken.AccessToken,
				Service:   controller.servicesService.FindByName(schemas.Google),
				UserId:    user.Id,
				User:      user,
				ServiceId: controller.servicesService.FindByName(schemas.Google).Id,
			})
			if err != nil {
				return "", err
			}
			newSessionToken, _ := controller.userService.Login(user, controller.servicesService.FindByName(schemas.Google))
			ctx.Redirect(http.StatusFound, "http://localhost:8081/callback?code="+codeCredentials.Code+"&state="+codeCredentials.State)
			return newSessionToken, nil
		}
	}
	// githubServ, tokenRepositoryice := controller.servicesService.FindByName(schemas.Github)
	// userInfo, err := controller.service.GetUserInfo(githubTokenResponse.AccessToken)
	serviceUserInfos := schemas.ServicesUserInfos{}
	userInfos := controller.servicesService.GetUserInfosByToken(googleServiceToken.AccessToken, schemas.Google)
	userInfos(&serviceUserInfos)
	fmt.Printf("serviceUserInfos: %++v\n", serviceUserInfos)
	// userInfo := ServicesUserInfos.GithubUserInfos
	return "", nil
}
