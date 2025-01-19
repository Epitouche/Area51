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

func (api *UserApi) GetServices(ctx *gin.Context) {
	if allServices, err := api.userController.GetAllServices(ctx); err != nil {
		ctx.JSON(http.StatusOK, []schemas.Service{})
	} else {
		ctx.JSON(http.StatusOK, allServices)
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

func (api *UserApi) LogoutService(ctx *gin.Context) {
	if err := api.userController.LogoutService(ctx); err != nil {
		ctx.JSON(http.StatusNotFound, &schemas.BasicResponse{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.BasicResponse{
			Message: "Successfully logged out",
		})
	}
}

func (api *UserApi) DeleteAccount(ctx *gin.Context) {
	if err := api.userController.DeleteAccount(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.BasicResponse{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.BasicResponse{
			Message: "Account successfully deleted",
		})
	}
}