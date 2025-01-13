package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
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
	if token, err := api.controller.StoreMobileToken(ctx); err != nil {
		ctx.JSON(http.StatusNotFound, schemas.BasicResponse{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	}
}
