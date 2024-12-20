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

type GithubController interface {
	RedirectionToGithubService(ctx *gin.Context, path string) (string, error)
	ServiceGithubCallback(ctx *gin.Context, path string) (string, error)
	GetUserInfos(ctx *gin.Context) (userInfos schemas.GithubUserInfo, err error)
	StoreMobileToken(ctx *gin.Context) (string, error)
}

type githubController struct {
	service         services.GithubService
	userService     services.UserService
	serviceToken    services.TokenService
	servicesService services.ServicesService
}

func NewGithubController(
	service services.GithubService,
	userService services.UserService,
	serviceToken services.TokenService,
	servicesService services.ServicesService,
) GithubController {
	return &githubController{
		service:         service,
		userService:     userService,
		serviceToken:    serviceToken,
		servicesService: servicesService,
	}
}

func (controller *githubController) RedirectionToGithubService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("GITHUB_CLIENT_ID")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	state, err := toolbox.GenerateCSRFToken()
	if err != nil {
		return "", err
	}

	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)
	redirectUri := fmt.Sprintf("%s%s/callback", appAdressHost, appPort)
	authUrl := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&response_type=code&scope=repo&redirect_uri=%s&state=%s", clientId, redirectUri, state)
	return authUrl, nil
}

func (controller *githubController) ServiceGithubCallback(ctx *gin.Context, path string) (string, error) {
	var codeCredentials schemas.GithubCodeCredentials

	err := json.NewDecoder(ctx.Request.Body).Decode(&codeCredentials)
	if err != nil || codeCredentials.Code == "" || codeCredentials.State == "" {
		return "", fmt.Errorf("unable to decode credentials because %w", err)
	}

	githubTokenResponse, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
	if err != nil {
		return "", fmt.Errorf("unable to get access token because %w", err)
	}

	githubService := controller.servicesService.FindByName(schemas.Github)
	userInfo, err := controller.service.GetUserInfo(githubTokenResponse.AccessToken)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}

	isAlreadyRegistered := false
	var actualUser schemas.User
	if userInfo.Email == "" {
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
	} else {
		actualUser = controller.userService.GetUserByEmail(userInfo.Email)
	}
	if actualUser.Email != "" || actualUser.Username != "" {
		isAlreadyRegistered = true
	} else {
		newUser := schemas.User{
			Username: userInfo.Login,
			Email: userInfo.Email,
		}
		if err := controller.userService.CreateUser(newUser); err == nil {
			actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		} else {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		isAlreadyRegistered = true
	}

	var newGithubToken schemas.ServiceToken
	var tokenId *uint64
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id:      *actualUser.TokenId,
			Token:   githubTokenResponse.AccessToken,
			Service: githubService,
			UserId:  actualUser.Id,
			User:    actualUser,
		}
		if actualUser.TokenId != nil {
			actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
			if actualServiceToken.Token != "" {
				err := controller.serviceToken.Update(newGithubToken)
				if err != nil {
					return "", fmt.Errorf("unable to update token because %w", err)
				}
				tokenId = &actualServiceToken.Id
			}
		}
	}

	if tokenId == nil {
		if savedTokenId, err := controller.serviceToken.SaveToken(newGithubToken); err != nil {
			return "", fmt.Errorf("unable to save token because %w", err)
		} else {
			tokenId = &savedTokenId
		}
	}

	actualUser.TokenId = tokenId
	var token string
	if isAlreadyRegistered {
		token, err = controller.userService.Login(actualUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		ctx.Redirect(http.StatusFound, fmt.Sprintf("http://localhost:8081/callback?code=%s&state=%s", codeCredentials.Code, codeCredentials.State))
	} else {
		token, err = controller.userService.Register(actualUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
	}
	return token, nil
}

func (controller *githubController) GetUserInfos(ctx *gin.Context) (userInfos schemas.GithubUserInfo, err error) {
	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	user, err := controller.userService.GetUserInfos(tokenString)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	token, err := controller.serviceToken.GetTokenById(*user.TokenId)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}

	githubUserInfos, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	return githubUserInfos, nil
}

func (controller *githubController) StoreMobileToken(ctx *gin.Context) (string, error) {
	var result schemas.MobileToken
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	githubService := controller.servicesService.FindByName(schemas.Github)
	userInfo, err := controller.service.GetUserInfo(result.Token)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}

	isAlreadyRegistered := false
	var actualUser schemas.User
	if userInfo.Email == "" {
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
	} else {
		actualUser = controller.userService.GetUserByEmail(userInfo.Email)
	}
	if actualUser.Email != "" || actualUser.Username != "" {
		isAlreadyRegistered = true
	} else {
		newUser := schemas.User{
			Username: userInfo.Login,
			Email: userInfo.Email,
		}
		if err := controller.userService.CreateUser(newUser); err == nil {
			actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		} else {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		isAlreadyRegistered = true
	}

	var newGithubToken schemas.ServiceToken
	var tokenId *uint64
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id:      *actualUser.TokenId,
			Token:   result.Token,
			Service: githubService,
			UserId:  actualUser.Id,
		}
		if actualUser.TokenId != nil {
			actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
			if actualServiceToken.Token != "" {
				err := controller.serviceToken.Update(newGithubToken)
				if err != nil {
					return "", fmt.Errorf("unable to update token because %w", err)
				}
				tokenId = &actualServiceToken.Id
			}
		}
	}

	if tokenId == nil {
		savedTokenId, _ := controller.serviceToken.SaveToken(newGithubToken)
		tokenId = &savedTokenId
	}

	actualUser.TokenId = tokenId
	var token string
	if isAlreadyRegistered {
		token, err = controller.userService.Login(actualUser)
		if err != nil {
			return "", fmt.Errorf("unable to login user because %w", err)
		}
	} else {
		token, err = controller.userService.Register(actualUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
	}
	return token, nil
}
