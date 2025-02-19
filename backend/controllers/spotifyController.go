package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"area51/database"
	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

type SpotifyController interface {
	RedirectionToSpotifyService(*gin.Context, string) (string, error)
	ServiceSpotifyCallback(*gin.Context, string) (string, error)
}

type spotifyController struct {
	service         services.SpotifyService
	servicesService services.ServicesService
	userService     services.UserService
	serviceToken    services.TokenService
}

func NewSpotifyController(
	service services.SpotifyService,
	servicesService services.ServicesService,
	userService services.UserService,
	serviceToken services.TokenService,
) SpotifyController {
	return &spotifyController{
		service:         service,
		servicesService: servicesService,
		userService:     userService,
		serviceToken:    serviceToken,
	}
}

func (controller *spotifyController) RedirectionToSpotifyService(ctx *gin.Context, path string) (string, error) {
	clientId := toolbox.GetInEnv("SPOTIFY_CLIENT_ID")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAddressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	state, err := toolbox.GenerateCSRFToken()
	if err != nil {
		return "", err
	}
	ctx.SetCookie("latestCSRFToken", state, 3600, "/", "localhost", false, true)
	redirectUri := fmt.Sprintf("%s%s/callback", appAddressHost, appPort)
	scopes := []string{
		"playlist-read-private",
		"playlist-modify-public",
		"playlist-modify-private",
		"user-read-private",
		"user-read-email",
	}
	scope := strings.Join(scopes, " ")
	authUrl := fmt.Sprintf(
		"https://accounts.spotify.com/authorize?client_id=%s&redirect_uri=%s&state=%s&response_type=code&scope=%s",
		clientId,
		redirectUri,
		state,
		url.QueryEscape(scope),
	)
	return authUrl, nil
}

func (controller *spotifyController) ServiceSpotifyCallback(ctx *gin.Context, path string) (string, error) {
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
	spotifyTokenResponse, err := controller.service.AuthGetServiceAccessToken(codeCredentials.Code, path)
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
				Token:     spotifyTokenResponse.AccessToken,
				Service:   controller.servicesService.FindByName(schemas.Spotify),
				UserId:    user.Id,
				User:      user,
				ServiceId: controller.servicesService.FindByName(schemas.Spotify).Id,
			})
			if err != nil {
				return "", err
			}
			newSessionToken, _ := controller.userService.Login(user, controller.servicesService.FindByName(schemas.Spotify))
			ctx.Redirect(http.StatusFound, "http://localhost:8081/callback?code="+codeCredentials.Code+"&state="+codeCredentials.State)
			return newSessionToken, nil
		}
	}

	spotifyService := controller.servicesService.FindByName(schemas.Spotify)
	var ServicesUserInfos schemas.ServicesUserInfos
	userInfos := controller.servicesService.GetUserInfosByToken(spotifyTokenResponse.AccessToken, schemas.Spotify)
	userInfos(&ServicesUserInfos)
	userInfo := ServicesUserInfos.SpotifyUserInfos
	var actualUser schemas.User
	actualUser = controller.userService.GetUserByEmail(&userInfo.Email)
	if actualUser.Email != nil {
		isAlreadyRegistered = true
	}

	var newSpotifyToken schemas.ServiceToken
	var newUser schemas.User
	password, err := database.HashPassword(toolbox.GetInEnv("DEFAULT_PASSWORD"))
	if err != nil {
		return "", fmt.Errorf("unable to hash password because %w", err)
	}
	serviceToken, _ := controller.userService.GetServiceByIdForUser(actualUser, spotifyService.Id)
	if isAlreadyRegistered {
		newSpotifyToken = schemas.ServiceToken{
			Id:        serviceToken.Id,
			Token:     spotifyTokenResponse.AccessToken,
			Service:   spotifyService,
			UserId:    actualUser.Id,
			User:      actualUser,
			ServiceId: controller.servicesService.FindByName(schemas.Spotify).Id,
		}
	} else {
		newUser = schemas.User{
			Username: userInfo.DisplayName,
			Email:    &userInfo.Email,
			Password: &password,
		}
		err = controller.userService.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("unable to create user because %w", err)
		}
		actualUser = controller.userService.GetUserByEmail(&userInfo.Email)

		newSpotifyToken = schemas.ServiceToken{
			Token:        spotifyTokenResponse.AccessToken,
			RefreshToken: spotifyTokenResponse.RefreshToken,
			Service:      spotifyService,
			UserId:       actualUser.Id,
			User:         actualUser,
			ServiceId:    controller.servicesService.FindByName(schemas.Spotify).Id,
		}
		err = controller.userService.AddServiceToUser(actualUser, newSpotifyToken)
		if err != nil {
			return "", fmt.Errorf("unable to add service to user because %w", err)
		}
		isAlreadyRegistered = true
	}
	if serviceToken.Id != 0 {
		actualServiceToken, _ := controller.serviceToken.GetTokenByUserIdAndServiceId(actualUser.Id, spotifyService.Id)
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
			Email:    &userInfo.Email,
			Password: &password,
		}
	} else {
		tokens, _ := controller.serviceToken.GetTokenByUserId(actualUser.Id)
		for _, token := range tokens {
			if token.UserId == actualUser.Id {
				newUser = schemas.User{
					Username: userInfo.DisplayName,
					Email:    &userInfo.Email,
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
		token, _ := controller.userService.Login(newUser, spotifyService)
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
