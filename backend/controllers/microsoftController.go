package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/database"
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
	serviceToken    services.TokenService
}

func NewMicrosoftController(
	service services.MicrosoftService,
	userService services.UserService,
	servicesService services.ServicesService,
	serviceToken services.TokenService,
) MicrosoftController {
	return &microsoftController{
		service:         service,
		userService:     userService,
		servicesService: servicesService,
		serviceToken:    serviceToken,
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
	redirectUri := appAdressHost + appPort + path
	authUrl := fmt.Sprintf("https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=%s&response_type=code&scope=openid profile Calendars.Read Calendars.ReadWrite Calendars.ReadWrite.Shared Calendars.Read.Shared Chat.Read Mail.Send https://graph.microsoft.com/User.Read&redirect_uri=%s&state=%s", clientId, redirectUri, state)
	return authUrl, nil
}

func (controller *microsoftController) ServiceMicrosoftCallback(ctx *gin.Context, path string) (string, error) {
	var isAlreadyRegistered bool = false
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
	microsoftService := controller.servicesService.FindByName(schemas.Microsoft)
	actualUserInfos := schemas.ServicesUserInfos{
		GithubUserInfos:    nil,
		SpotifyUserInfos:   nil,
		MicrosoftUserInfos: nil,
	}
	userInfos := controller.servicesService.GetUserInfosByToken(microsoftTokenResponse.AccessToken, schemas.Microsoft)
	userInfos(&actualUserInfos)
	userInfo := actualUserInfos.MicrosoftUserInfos
	var actualUser schemas.User
	actualUser = controller.userService.GetUserByEmail(&userInfo.Mail)
	if actualUser.Email != nil {
		isAlreadyRegistered = true
	}

	var newSpotifyToken schemas.ServiceToken
	var newUser schemas.User
	password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
	if err != nil {
		return "", fmt.Errorf("unable to hash password because %w", err)
	}
	serviceToken, _ := controller.userService.GetServiceByIdForUser(actualUser, microsoftService.Id)
	if isAlreadyRegistered {
		newSpotifyToken = schemas.ServiceToken{
			Id:        serviceToken.Id,
			Token:     microsoftTokenResponse.AccessToken,
			Service:   microsoftService,
			ServiceId: controller.servicesService.FindByName(schemas.Microsoft).Id,
			UserId:    actualUser.Id,
			User:      actualUser,
		}
	} else {
		newUser = schemas.User{
			Username: userInfo.DisplayName,
			Email:    &userInfo.Mail,
			Password: &password,
		}
		err = controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByEmail(&userInfo.Mail)

		newSpotifyToken = schemas.ServiceToken{
			Token:        microsoftTokenResponse.AccessToken,
			RefreshToken: microsoftTokenResponse.RefreshToken,
			Service:      microsoftService,
			ServiceId:    controller.servicesService.FindByName(schemas.Microsoft).Id,
			UserId:       actualUser.Id,
			User:         actualUser,
		}
		err = controller.userService.AddServiceToUser(actualUser, newSpotifyToken)
		if err != nil {
			return "", fmt.Errorf("unable to add service to user because %w", err)
		}
		isAlreadyRegistered = true
	}
	if serviceToken.Id != 0 {
		actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, microsoftService.Id)
		if actualServiceToken.Token != "" {
			err := controller.serviceToken.Update(newSpotifyToken)
			if err != nil {
				return "", fmt.Errorf("unable to update token because %w", err)
			}
		}
	}

	if newUser.Username == "" {

		newUser = schemas.User{
			Username: userInfo.DisplayName,
			Email:    &userInfo.Mail,
			Password: &password,
		}
	} else {
		tokens, _ := controller.serviceToken.GetTokenByUserId(actualUser.Id)
		for _, token := range tokens {
			if token.UserId == actualUser.Id {
				newUser = schemas.User{
					Username: userInfo.DisplayName,
					Email:    &userInfo.Mail,
					Password: &password,
				}
				serviceToken.Id = token.Id
				err := controller.userService.UpdateUserInfos(actualUser)
				if err != nil {
					return "", fmt.Errorf("unable to update user infos because %w", err)
				}
				break
			}
		}
	}

	if isAlreadyRegistered {
		token, _ := controller.userService.Login(newUser, microsoftService)
		ctx.Redirect(http.StatusFound, "http://localhost:8081/callback?code="+codeCredentials.Code+"&state="+codeCredentials.State)
		return token, nil
	} else {
		token, err := controller.userService.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		return token, nil
	}

}
