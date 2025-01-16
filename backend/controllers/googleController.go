package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"area51/database"
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
	serviceToken    services.TokenService
}

func NewGoogleController(
	service services.GoogleService,
	userService services.UserService,
	servicesService services.ServicesService,
	serviceToken services.TokenService,
) GoogleController {
	return &googleController{
		service:         service,
		userService:     userService,
		servicesService: servicesService,
		serviceToken:    serviceToken,
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

	scopes := "openid https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/gmail.readonly https://www.googleapis.com/auth/gmail.labels https://www.googleapis.com/auth/gmail.modify https://www.googleapis.com/auth/gmail.metadata"

	authUrl := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=%s&redirect_uri=%s&state=%s",
		clientId,
		url.QueryEscape(scopes),
		url.QueryEscape(redirectUri),
		state,
	)
	return authUrl, nil
}

func (controller *googleController) ServiceGoogleCallback(ctx *gin.Context, path string) (string, error) {
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
	googleServiceToken, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
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
	googleService := controller.servicesService.FindByName(schemas.Google)
	// userInfo, err := controller.service.GetUserInfo(githubTokenResponse.AccessToken)
	serviceUserInfos := schemas.ServicesUserInfos{}
	userInfos := controller.servicesService.GetUserInfosByToken(googleServiceToken.AccessToken, schemas.Google)
	userInfos(&serviceUserInfos)
	userInfo := serviceUserInfos.GoogleUserInfos

	var actualUser schemas.User
	actualUser = controller.userService.GetUserByEmail(&userInfo.Email)
	if actualUser.Email != nil {
		isAlreadyRegistered = true
	}

	var newGoogleToken schemas.ServiceToken
	var newUser schemas.User
	password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
	if err != nil {
		return "", fmt.Errorf("unable to hash password because %w", err)
	}
	serviceToken, _ := controller.userService.GetServiceByIdForUser(actualUser, googleService.Id)
	if isAlreadyRegistered {
		newGoogleToken = schemas.ServiceToken{
			Id:        serviceToken.Id,
			Token:     googleServiceToken.AccessToken,
			Service:   googleService,
			UserId:    actualUser.Id,
			User:      actualUser,
			ServiceId: controller.servicesService.FindByName(schemas.Google).Id,
		}
	} else {
		var email *string
		if userInfo.Email == "" {
			email = nil
		} else {
			email = &userInfo.Email
		}
		newUser = schemas.User{
			Username: userInfo.Name,
			Email:    email,
			Password: &password,
		}
		err = controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByEmail(&userInfo.Email)
		newGoogleToken = schemas.ServiceToken{
			Token:        googleServiceToken.AccessToken,
			RefreshToken: googleServiceToken.RefreshToken,
			Service:      googleService,
			UserId:       actualUser.Id,
			User:         actualUser,
			ServiceId:    controller.servicesService.FindByName(schemas.Google).Id,
		}
		err = controller.userService.AddServiceToUser(actualUser, newGoogleToken)
		if err != nil {
			return "", fmt.Errorf("unable to add service to user because %w", err)
		}
		isAlreadyRegistered = true
	}
	if serviceToken.Id != 0 {
		actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, googleService.Id)
		if actualServiceToken.Token != "" {
			err := controller.serviceToken.Update(newGoogleToken)
			if err != nil {
				return "", fmt.Errorf("unable to update token because %w", err)
			}
		}
	}

	if newUser.Username == "" {
		var email *string
		if userInfo.Email == "" {
			email = nil
		} else {
			email = &userInfo.Email
		}
		password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
		if err != nil {
			return "", fmt.Errorf("unable to hash password because %w", err)
		}
		newUser = schemas.User{
			Username: userInfo.Name,
			Email:    email,
			Password: &password,
		}
	} else {
		tokens, _ := controller.serviceToken.GetTokenByUserId(actualUser.Id)
		for _, token := range tokens {
			if token.UserId == actualUser.Id {
				var email *string
				if userInfo.Email == "" {
					email = nil
				} else {
					email = &userInfo.Email
				}
				newUser = schemas.User{
					Username: userInfo.Name,
					Email:    email,
					Password: &password,
				}
				serviceToken.Id = token.Id
				err = controller.userService.UpdateUserInfos(actualUser)
				if err != nil {
					return "", fmt.Errorf("unable to update user infos because %w", err)
				}
				break
			}
		}
	}
	if isAlreadyRegistered {
		token, _ := controller.userService.Login(newUser, googleService)
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
