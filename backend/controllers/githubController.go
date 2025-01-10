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

type GithubController interface {
	RedirectionToGithubService(ctx *gin.Context, path string) (string, error)
	ServiceGithubCallback(ctx *gin.Context, path string) (string, error)
	GetUserInfos(ctx *gin.Context, serviceName schemas.ServiceName) (userInfos schemas.GithubUserInfo, err error)
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

	var isAlreadyRegistered bool = false
	var codeCredentials schemas.GithubCodeCredentials
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
	githubTokenResponse, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
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
			controller.userService.AddServiceToUser(user, schemas.ServiceToken{
				Token:   githubTokenResponse.AccessToken,
				Service: controller.servicesService.FindByName(schemas.Github),
				UserId:  user.Id,
			})
			newSessionToken, _ := controller.userService.Login(user, controller.servicesService.FindByName(schemas.Github))
			ctx.Redirect(http.StatusFound, "http://localhost:8081/callback?code="+codeCredentials.Code+"&state="+codeCredentials.State)
			return newSessionToken, nil
		}
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
		actualUser = controller.userService.GetUserByEmail(&userInfo.Email)
	}
	if actualUser.Email != nil {
		isAlreadyRegistered = true
	}

	var newGithubToken schemas.ServiceToken
	var newUser schemas.User
	var tokenId *uint64
	serviceToken, _ := controller.userService.GetServiceByIdForUser(actualUser, githubService.Id)
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id:      serviceToken.Id,
			Token:   githubTokenResponse.AccessToken,
			Service: githubService,
			UserId:  actualUser.Id,
			User:    actualUser,
		}
		if serviceToken.Id != 0 {
			actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
			if actualServiceToken.Token != "" {
				err := controller.serviceToken.Update(newGithubToken)
				if err != nil {
					return "", fmt.Errorf("unable to update token because %w", err)
				}
				tokenId = &actualServiceToken.Id
			} else {
				tokenId = nil
			}
		}
	} else {
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
			Username: userInfo.Login,
			Email:    email,
			Password: &password,
		}
		err = controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByUsername(userInfo.Login)
		newGithubToken = schemas.ServiceToken{
			Token:        githubTokenResponse.AccessToken,
			RefreshToken: githubTokenResponse.RefreshToken,
			Service:      githubService,
			UserId:       actualUser.Id,
			User:         actualUser,
		}
		controller.userService.AddServiceToUser(actualUser, newGithubToken)
		isAlreadyRegistered = true
	}

	if tokenId == nil {
		_, err := controller.serviceToken.SaveToken(newGithubToken)
		if err != nil {
			return "", fmt.Errorf("unable to save token because %w", err)
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
			Username: userInfo.Login,
			Email:    email,
			Password: &password,
			// TokenId:  tokenId,
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
				password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
				if err != nil {
					return "", fmt.Errorf("unable to hash password because %w", err)
				}
				newUser = schemas.User{
					Username: userInfo.Login,
					Email:    email,
					Password: &password,
					// TokenId:  &token.Id,
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

func (controller *githubController) GetUserInfos(ctx *gin.Context, serviceName schemas.ServiceName) (userInfos schemas.GithubUserInfo, err error) {
	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	user, err := controller.userService.GetUserInfos(tokenString)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	token, err := controller.serviceToken.GetTokenByUserId(user.Id)
	// token, err := controller.serviceToken.GetTokenById(*user.TokenId)
	if err != nil {
		return schemas.GithubUserInfo{}, err
	}
	for _, actualToken := range token {
		if actualToken.ServiceId == controller.servicesService.FindByName(serviceName).Id {
			githubUserInfos, err := controller.service.GetUserInfo(actualToken.Token)
			if err != nil {
				return schemas.GithubUserInfo{}, err
			}
			return githubUserInfos, nil
		}
	}
	return schemas.GithubUserInfo{}, nil
}

func (controller *githubController) StoreMobileToken(ctx *gin.Context) (string, error) {
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
		actualUser = controller.userService.GetUserByEmail(&userInfo.Email)
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
			controller.userService.AddServiceToUser(user, schemas.ServiceToken{
				Token:   result.Token,
				Service: controller.servicesService.FindByName(result.Service),
				UserId:  user.Id,
			})
			newSessionToken, _ := controller.userService.Login(user, controller.servicesService.FindByName(result.Service))
			return newSessionToken, nil
		}
	}
	var newGithubToken schemas.ServiceToken
	var newUser schemas.User
	var tokenId *uint64
	serviceToken, _ := controller.userService.GetServiceByIdForUser(actualUser, githubService.Id)
	if isAlreadyRegistered {
		newGithubToken = schemas.ServiceToken{
			Id:      serviceToken.Id,
			Token:   result.Token,
			Service: githubService,
			UserId:  actualUser.Id,
		}
		if serviceToken.Id != 0 {
			actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, githubService.Id)
			if actualServiceToken.Token != "" {
				err := controller.serviceToken.Update(newGithubToken)
				if err != nil {
					return "", fmt.Errorf("unable to update token because %w", err)
				}
				tokenId = &actualServiceToken.Id
			}
		}
	} else {
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
			Username: userInfo.Login,
			Email:    email,
			Password: &password,
		}
		err = controller.userService.CreateUser(newUser)
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
		controller.userService.AddServiceToUser(actualUser, newGithubToken)
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
			Username: userInfo.Login,
			Email:    email,
			Password: &password,
			// TokenId:  tokenId,
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
				password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
				if err != nil {
					return "", fmt.Errorf("unable to hash password because %w", err)
				}
				newUser = schemas.User{
					Username: userInfo.Login,
					Email:    email,
					Password: &password,
					// TokenId:  &token.Id,
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
