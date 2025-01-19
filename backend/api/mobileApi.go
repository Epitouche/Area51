package api

import (
	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/toolbox"
)

type MobileApi struct {
	controller controllers.MobileController
}

func NewMobileApi(controller controllers.MobileController) *MobileApi {
	return &MobileApi{
		controller: controller,
	}
}

func (api *MobileApi) StoreMobileToken(ctx *gin.Context) {
	token, err := api.controller.StoreMobileToken(ctx)
	toolbox.HandleError(ctx, err, gin.H{"token": token})
}
