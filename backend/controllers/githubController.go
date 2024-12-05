package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

type GithubController interface {
	RedirectionToGithubService(ctx *gin.Context, path string) (string, error)
	serviceCallback(ctx *gin.Context, path string) (string, error)
	GetUserInfos(ctx *gin.Context) (userInfos schemas.GithubUserInfo, err error)
}

type githubController struct {
	service 	services.GithubService
	userService services.UserService
	serviceToken services.TokenService
}

func NewGithubController(
	service services.GithubService,
	userService services.UserService,
	serviceToken services.TokenService,
) GithubController {
	return &githubController{
		service: services.NewGithubService(),
		userService: services.NewUserService(),
		serviceToken: services.NewTokenService(),
	}
}


func (controller *githubController) RedirectionToGithubService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("GITHUB_CLIENT_ID")
	appPort := toolbox.GetInEnv("APP_PORT")
	appAdressHost := toolbox.GetInEnv("APP_ADRESS_HOST")

	state, err := toolbox.GenerateCSRFToken()
	if err != nil {
		return "", err
	}

	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)

	redirectUri := appAdressHost + appPort + path
	authUrl := "https://github.com/login/oauth/authorize" +
		"?client_id=" + clientId +
		"&response_type=code" +
		"&scope=repo" +
		"&redirect_uri=" + redirectUri +
		"&state=" + state
	return authUrl, nil
}

func (controller *githubController) serviceCallback(ctx *gin.Context, path string) (string, error) {
	code := ctx.Query("code")
	if code == "" {
		return "", nil
	}
	state := ctx.Query("state")
	latestCSRFToken, err := ctx.Cookie("latestCSRFToken")
	if state == "" {
		return "", nil
	}
	if state != latestCSRFToken {
		return "", nil
	}
	githubTokenResponse, err := controller.service.AuthGetServiceAccessToken(code, path)
	if err != nil {
		return "", err
	}
	newGithubToken := schemas.ServiceToken{
		Token:   githubTokenResponse.AccessToken,
		// Service: githubController.service,
		UserId:  1,
	}
	controller.serviceToken.SaveToken(newGithubToken)

	userAlreadExists := false
	if err != nil {
		if err.Error() == "token already exists" {
			userAlreadExists = true
		} else {
			return "", fmt.Errorf("unable to save token because %w", err)
		}
	}

	userInfo, err := controller.service.GetUserInfo(newGithubToken.Token)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}

	newUser := schemas.User{
		Username: userInfo.Login,
		Email:    userInfo.Email,
		// TokenId:  token,
	}

	if userAlreadExists {
		token, err := controller.userService.Login(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to login user because %w", err)
		}
		return token, nil
	} else {
		token, err := controller.userService.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		return token, nil
	}
}

func (controller *githubController) GetUserInfos(ctx *gin.Context) (userInfos schemas.GithubUserInfo, err error) {
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err := controller.userService.

}