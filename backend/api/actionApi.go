package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
)

type ActionApi struct {
	actionController controllers.ActionController
}

func NewActionApi(controller controllers.ActionController) *ActionApi {
	return &ActionApi{
		actionController: controller,
	}
}

func (api *ActionApi) CreateAction(ctx *gin.Context) {
	if message, err := api.actionController.CreateAction(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, message)
	}
}
