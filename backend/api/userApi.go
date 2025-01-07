package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
)

type UserDependencies struct {
	UserApi *UserApi
}

type UserApi struct {
	userController controllers.UserController
}

func NewUserApi(controller controllers.UserController) *UserApi {
	return &UserApi{
		userController: controller,
	}
}

func (api *UserApi) Login(ctx *gin.Context) {
	if token, err := api.userController.Login(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, &schemas.BasicResponse{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.JWT{
			Token: token,
		})
	}
}

func (api *UserApi) Register(ctx *gin.Context) {
	if token, err := api.userController.Register(ctx); err != nil {
		ctx.JSON(http.StatusConflict, &schemas.BasicResponse{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.JWT{
			Token: token,
		})
	}
}

func (api *UserApi) GetAccessToken(ctx *gin.Context) {
	if token, err := api.userController.GetAccessToken(ctx); err != nil {
		ctx.JSON(http.StatusUnauthorized, &schemas.BasicResponse{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.JWT{
			Token: token,
		})
	}
}

func (api *UserApi) GetServices(ctx *gin.Context) {
	allServices, err := api.userController.GetAllServices(ctx)
	if allServices == nil {
		ctx.JSON(http.StatusOK, []schemas.Service{})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
}


func (api *UserApi) GetWorkflows(ctx *gin.Context) {
	allWorkflows, err := api.userController.GetAllWorkflows(ctx)
	if allWorkflows == nil {
		ctx.JSON(http.StatusOK, []schemas.WorkflowJson{})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, allWorkflows)
}