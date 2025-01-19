package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area51/database"
	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

type MobileController interface {
	StoreMobileToken(ctx *gin.Context) (string, error)
}

type mobileController struct {
	userService     services.UserService
	serviceToken    services.TokenService
	servicesService services.ServicesService
}

func NewMobileController(
	userService services.UserService,
	serviceToken services.TokenService,
	servicesService services.ServicesService,
) MobileController {
	return &mobileController{
		userService:     userService,
		serviceToken:    serviceToken,
		servicesService: servicesService,
	}
}

func (controller *mobileController) StoreMobileToken(ctx *gin.Context) (string, error) {
	var result schemas.MobileToken
	var isAlreadyRegistered bool = false
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	githubService := controller.servicesService.FindByName(result.Service)
	if githubService == (schemas.Service{}) {
		return "", fmt.Errorf("service %s not found", result.Service)
	}
	var servicesUserInfos schemas.ServicesUserInfos
	userInfos := controller.servicesService.GetUserInfosByToken(result.Token, result.Service)
	userInfos(&servicesUserInfos)
	var infos schemas.MobileUsefulInfos
	switch result.Service {
	case schemas.Github:
		infos = schemas.MobileUsefulInfos{
			Login: servicesUserInfos.GithubUserInfos.Login,
			Email: servicesUserInfos.GithubUserInfos.Email,
		}
	case schemas.Spotify:
		infos = schemas.MobileUsefulInfos{
			Login: servicesUserInfos.SpotifyUserInfos.DisplayName,
			Email: servicesUserInfos.SpotifyUserInfos.Email,
		}
	case schemas.Google:
		infos = schemas.MobileUsefulInfos{
			Login: servicesUserInfos.GoogleUserInfos.Name,
			Email: servicesUserInfos.GoogleUserInfos.Email,
		}
	case schemas.Microsoft:
		infos = schemas.MobileUsefulInfos{
			Login: servicesUserInfos.MicrosoftUserInfos.DisplayName,
			Email: servicesUserInfos.MicrosoftUserInfos.Mail,
		}
	}

	var actualUser schemas.User
	if infos.Email == "" {
		actualUser = controller.userService.GetUserByUsername(infos.Login)
		if actualUser.Username != "" {
			isAlreadyRegistered = true
		}
	}
	if infos.Email != "" {
		actualUser = controller.userService.GetUserByEmail(&infos.Email)
	}
	if actualUser.Email != nil {
		isAlreadyRegistered = true
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
				Token:     result.Token,
				Service:   controller.servicesService.FindByName(result.Service),
				UserId:    user.Id,
				User: 	user,
				ServiceId: controller.servicesService.FindByName(result.Service).Id,
			})
			if err != nil {
				return "", err
			}
			newSessionToken, _ := controller.userService.Login(user, controller.servicesService.FindByName(result.Service))
			return newSessionToken, nil
		}
	}
	var newGithubToken schemas.ServiceToken
	var newUser schemas.User
	password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
	if err != nil {
		return "", fmt.Errorf("unable to hash password because %w", err)
	}
	serviceToken, _ := controller.userService.GetServiceByIdForUser(actualUser, githubService.Id)
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id:      serviceToken.Id,
			Token:   result.Token,
			Service: githubService,
			UserId:  actualUser.Id,
		}
	} else {
		var email *string
		if infos.Email == "" {
			email = nil
		} else {
			email = &infos.Email
		}

		newUser = schemas.User{
			Username: infos.Login,
			Email:    email,
			Password: &password,
		}
		err = controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByUsername(infos.Login)
		newGithubToken = schemas.ServiceToken{
			Id:        serviceToken.Id,
			Token:     result.Token,
			Service:   githubService,
			UserId:    actualUser.Id,
			User:      actualUser,
			ServiceId: controller.servicesService.FindByName(result.Service).Id,
		}
		err = controller.userService.AddServiceToUser(actualUser, newGithubToken)
		if err != nil {
			return "", fmt.Errorf("unable to add service to user because %w", err)
		}
		isAlreadyRegistered = true
	}
	if serviceToken.Id != 0 {
		actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
		if actualServiceToken.Token != "" {
			err := controller.serviceToken.Update(newGithubToken)
			if err != nil {
				return "", fmt.Errorf("unable to update token because %w", err)
			}
		}
	}

	if newUser.Username == "" {
		var email *string
		if infos.Email == "" {
			email = nil
		} else {
			email = &infos.Email
		}
		newUser = schemas.User{
			Username: infos.Login,
			Email:    email,
			Password: &password,
		}
	} else {
		tokens, _ := controller.serviceToken.GetTokenByUserId(actualUser.Id)
		for _, token := range tokens {
			if token.UserId == actualUser.Id {
				var email *string
				if infos.Email == "" {
					email = nil
				} else {
					email = &infos.Email
				}
				newUser = schemas.User{
					Username: infos.Login,
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
		token, _ := controller.userService.Login(newUser, githubService)
		return token, nil
	} else {
		token, err := controller.userService.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		return token, nil
	}
}
