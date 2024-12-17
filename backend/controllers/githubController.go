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

func (controller *githubController) ServiceGithubCallback(ctx *gin.Context, path string) (string, error) {
	var isAlreadyRegistered bool = false
	var codeCredentials schemas.GithubCodeCredentials
	err := json.NewDecoder(ctx.Request.Body).Decode(&codeCredentials)
	if err != nil {
		return "", err
	}
	// code := ctx.Query("code")
	if codeCredentials.Code == "" {
		return "", nil
	}
	// state := ctx.Query("state")
	// latestCSRFToken, err := ctx.Cookie("latestCSRFToken")
	if codeCredentials.State == "" {
		return "", nil
	}
	// if state != latestCSRFToken {
	// 	return "", nil
	// }
	githubTokenResponse, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
	if err != nil {
		return "", err
	}

	githubService := controller.servicesService.FindByName(schemas.Github)

	userInfo, err := controller.service.GetUserInfo(githubTokenResponse.AccessToken)

	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}
	var actualUser schemas.User
	if userInfo.Email == "" {
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		if actualUser.Username != "" {
			isAlreadyRegistered = true
		}
	}
	if userInfo.Email != "" {
		actualUser = controller.userService.GetUserByEmail(userInfo.Email)
	}
	if actualUser.Email != "" {
		isAlreadyRegistered = true
	}

	var newGithubToken schemas.ServiceToken
	var newUser schemas.User
	var tokenId *uint64
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id: 	*actualUser.TokenId,
			Token:   githubTokenResponse.AccessToken,
			Service: githubService,
			UserId:  actualUser.Id,
			User: actualUser,
		}
		if actualUser.TokenId != nil {
			actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
			if actualServiceToken.Token != "" {
				controller.serviceToken.Update(newGithubToken)
				tokenId = &actualServiceToken.Id
			}
		}
	} else {
		newUser = schemas.User{
			Username: userInfo.Login,
			Email:    userInfo.Email,
		}
		err := controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		newGithubToken = schemas.ServiceToken{
			Token:   githubTokenResponse.AccessToken,
			RefreshToken: githubTokenResponse.RefreshToken,
			Service: githubService,
			UserId:  actualUser.Id,
			User: actualUser,
		}
		isAlreadyRegistered = true
	}

	if tokenId == nil {
		savedTokenId, _ := controller.serviceToken.SaveToken(newGithubToken)
		tokenId = &savedTokenId
	}

	if newUser.Username == "" {
		newUser = schemas.User{
			Username: userInfo.Login,
			Email:    userInfo.Email,
			TokenId:  tokenId,
		}
	} else {
		tokens, _ := controller.serviceToken.GetTokenByUserId(actualUser.Id)
		for _, token := range tokens {
			if token.UserId == actualUser.Id {
				newUser = schemas.User{
					Username: userInfo.Login,
					Email:    userInfo.Email,
					TokenId: &token.Id,
				}
				actualUser.TokenId = &token.Id
				err := controller.userService.UpdateUserInfos(actualUser)
				if err != nil {
					return "", fmt.Errorf("unable to update user infos because %w", err)
				}
				break
			}
		}

	}

	if isAlreadyRegistered {
		token, _ := controller.userService.Login(newUser)
		// ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
		ctx.Redirect(http.StatusFound, "http://localhost:8081/callback?code=" + codeCredentials.Code + "&state=" + codeCredentials.State)
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
	var isAlreadyRegistered bool = false
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	githubService := controller.servicesService.FindByName(schemas.Github)

	userInfo, err := controller.service.GetUserInfo(result.Token)
	if err != nil {
		return "", fmt.Errorf("unable to get user info because %w", err)
	}
	var actualUser schemas.User
	if userInfo.Email == "" {
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		if actualUser.Username != "" {
			isAlreadyRegistered = true
		}
	}
	if userInfo.Email != "" {
		actualUser = controller.userService.GetUserByEmail(userInfo.Email)
	}
	if actualUser.Email != "" {
		isAlreadyRegistered = true
	}
	var newGithubToken schemas.ServiceToken
	var newUser schemas.User
	var tokenId *uint64
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id: 	*actualUser.TokenId,
			Token:   result.Token,
			Service: githubService,
			UserId:  actualUser.Id,
		}
		if actualUser.TokenId != nil {
			actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
			if actualServiceToken.Token != "" {
				controller.serviceToken.Update(newGithubToken)
				tokenId = &actualServiceToken.Id
			}
		}
	} else {
		newUser = schemas.User{
			Username: userInfo.Login,
			Email:    userInfo.Email,
		}
		err := controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		newGithubToken = schemas.ServiceToken{
			Token:   result.Token,
			Service: githubService,
			UserId:  actualUser.Id,
		}
		isAlreadyRegistered = true
	}

	if tokenId == nil {
		savedTokenId, _ := controller.serviceToken.SaveToken(newGithubToken)
		tokenId = &savedTokenId
	}

	if newUser.Username == "" {
		newUser = schemas.User{
			Username: userInfo.Login,
			Email:    userInfo.Email,
			TokenId:  tokenId,
		}
	} else {
		tokens, _ := controller.serviceToken.GetTokenByUserId(actualUser.Id)
		for _, token := range tokens {
			if token.UserId == actualUser.Id {
				newUser = schemas.User{
					Username: userInfo.Login,
					Email:    userInfo.Email,
					TokenId: &token.Id,
				}
				actualUser.TokenId = &token.Id
				err := controller.userService.UpdateUserInfos(actualUser)
				if err != nil {
					return "", fmt.Errorf("unable to update user infos because %w", err)
				}
				break
			}
		}

	}

	if isAlreadyRegistered {
		token, _ := controller.userService.Login(newUser)
		return token, nil
	} else {
		token, err := controller.userService.Register(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to register user because %w", err)
		}
		return token, nil
	}
}
