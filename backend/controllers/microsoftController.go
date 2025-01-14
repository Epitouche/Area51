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

type MicrosoftController interface {
	RedirectionToMicrosoftService(ctx *gin.Context, path string) (string, error)
	ServiceMicrosoftCallback(ctx *gin.Context, path string) (string, error)
}

type microsoftController struct {
	service         services.MicrosoftService
	userService     services.UserService
	servicesService services.ServicesService
}

func NewMicrosoftController(
	service services.MicrosoftService,
	userService services.UserService,
	servicesService services.ServicesService,
) MicrosoftController {
	return &microsoftController{
		service:         service,
		userService:     userService,
		servicesService: servicesService,
	}
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
	authUrl := fmt.Sprintf("https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=%s&response_type=code&scope=Mail.ReadWrite Mail.Read User.Read Mail.Send offline_access calendars.Read calendars.ReadWrite&redirect_uri=%s&state=%s", clientId, redirectUri, state)
	return authUrl, nil
}

func (controller *microsoftController) ServiceMicrosoftCallback(ctx *gin.Context, path string) (string, error) {
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
	microsoftTokenResponse, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
	if err != nil {
		return "", err
	}
	authHeader := ctx.GetHeader("Authorization")

	if authHeader != "" && len(authHeader) >= len("Bearer ") {
		fmt.Printf("JE SUIS DANS LE IF CE QUI N'A AUCUN SENS\n")
		token := authHeader[len("Bearer "):]
		user, err := controller.userService.GetUserInfos(token)
		if err != nil {
			return "", err
		}
		if user.Username != "" {
			err := controller.userService.AddServiceToUser(user, schemas.ServiceToken{
				Token:     microsoftTokenResponse.AccessToken,
				Service:   controller.servicesService.FindByName(schemas.Microsoft),
				UserId:    user.Id,
				User:      user,
				ServiceId: controller.servicesService.FindByName(schemas.Microsoft).Id,
			})
			if err != nil {
				return "", err
			}
			newSessionToken, _ := controller.userService.Login(user, controller.servicesService.FindByName(schemas.Microsoft))
			ctx.Redirect(http.StatusFound, "http://localhost:8081/callback?code="+codeCredentials.Code+"&state="+codeCredentials.State)
			return newSessionToken, nil
		}
	}
	// githubService := controller.servicesService.FindByName(schemas.Github)
	// userInfo, err := controller.service.GetUserInfo(githubTokenResponse.AccessToken)
	actualUserInfos := schemas.ServicesUserInfos{
		GithubUserInfos:    nil,
		SpotifyUserInfos:   nil,
		MicrosoftUserInfos: nil,
	}
	userInfos := controller.servicesService.GetUserInfosByToken(microsoftTokenResponse.AccessToken, schemas.Microsoft)
	userInfos(&actualUserInfos)
	userInfo := actualUserInfos
	fmt.Printf("UserInfos: %+v\n", userInfo)
	return "", nil
}
