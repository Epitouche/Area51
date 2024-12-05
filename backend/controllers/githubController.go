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
	ServiceCallback(ctx *gin.Context, path string) (string, error)
	GetUserInfos(ctx *gin.Context) (userInfos schemas.GithubUserInfo, err error)
}

type githubController struct {
	service 		services.GithubService
	userService 	services.UserService
	serviceToken 	services.TokenService
	servicesService services.ServicesService
}

func NewGithubController(
	service services.GithubService,
	userService services.UserService,
	serviceToken services.TokenService,
	servicesService services.ServicesService,
) GithubController {
	return &githubController{
		service: service,
		userService: userService,
		serviceToken: serviceToken,
		servicesService: servicesService,
	}
}


func (controller *githubController) RedirectionToGithubService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("GITHUB_CLIENT_ID")
	appPort := toolbox.GetInEnv("APP_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

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

func (controller *githubController) ServiceCallback(ctx *gin.Context, path string) (string, error) {
	fmt.Printf("I enter in the Controller!!!!!!!!!\n")
	code := ctx.Query("code")
	if code == "" {
		return "", nil
	}
	state := ctx.Query("state")
	// latestCSRFToken, err := ctx.Cookie("latestCSRFToken")
	if state == "" {
		return "", nil
	}
	fmt.Printf("state Passed !!!!!!!!!\n")
	// if state != latestCSRFToken {
	// 	return "", nil
	// }
	fmt.Printf("CSRF token passed !!!!!!!!!")
	githubTokenResponse, err := controller.service.AuthGetServiceAccessToken(code, path)
	fmt.Printf("githubTokenResponse Passed !!!!!!!!!")
	if err != nil {
		return "", err
	}

	githubService := controller.servicesService.FindByName(schemas.Github)
	fmt.Printf("githubService Passed !!!!!!!!!")

	newGithubToken := schemas.ServiceToken{
		Token:   githubTokenResponse.AccessToken,
		Service: githubService,
		UserId:  1,
	}
	fmt.Printf("newGithubToken Passed !!!!!!!!!")
	tokenId, err := controller.serviceToken.SaveToken(newGithubToken)
	userAlreadExists := false

	fmt.Printf("MON BOEUF\n")

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
		TokenId:  tokenId,
	}
	fmt.Printf("newUser username: %v\n", newUser.Username)
	fmt.Printf("newUser email: %v\n", newUser.Email)
	fmt.Printf("newUser tokenId: %v\n", newUser.TokenId)


	if userAlreadExists {
		fmt.Printf("user already exists")
		token, err := controller.userService.Login(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to login user because %w", err)
		}
		return token, nil
	} else {
		fmt.Printf("Creating new user")

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

	user, err := controller.userService.GetUserInfos(tokenString)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	token, err := controller.serviceToken.GetTokenById(user.TokenId)

	githubUserInfos, err := controller.service.GetUserInfo(token.Token)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	return githubUserInfos, nil
}