package api

import (
	"area51/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
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
	message, err := api.actionController.CreateAction(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, message)
}