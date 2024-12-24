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
	token, err := api.userController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &schemas.BasicResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &schemas.JWT{
		Token: token,
	})
}

func (api *UserApi) Register(ctx *gin.Context) {
	token, err := api.userController.Register(ctx)
	if err != nil {
		ctx.JSON(http.StatusConflict, &schemas.BasicResponse{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &schemas.JWT{
		Token: token,
	})
}

func (api *UserApi) GetServices(ctx *gin.Context) {
	allServices, err := api.userController.GetAllServices(ctx)
	if err != nil {
	}
	ctx.JSON(http.StatusOK, allServices)
}
