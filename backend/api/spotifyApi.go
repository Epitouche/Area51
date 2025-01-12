package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
)

type SpotifyApi struct {
	controller controllers.SpotifyController
}

func NewSpotifyApi(controller controllers.SpotifyController) *SpotifyApi {
	return &SpotifyApi{
		controller: controller,
	}
}

func (api *SpotifyApi) RedirectToSpotify(ctx *gin.Context, path string) {
	if authURL, err := api.controller.RedirectionToSpotifyService(ctx, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"spotify_authentication_url": authURL})
	}
}

func (api *SpotifyApi) HandleSpotifyTokenCallback(ctx *gin.Context, path string) {
	if spotify_token, err := api.controller.ServiceSpotifyCallback(ctx, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"access_token": spotify_token})
	}
}

func (api *SpotifyApi) StoreMobileToken(ctx *gin.Context) {
	if token, err := api.controller.StoreMobileToken(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}