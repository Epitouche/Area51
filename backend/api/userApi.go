package api

import (
	"area51/controllers"
	"area51/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userController controllers.UserController
}

func NewUserAPI(userController controllers.UserController) *UserApi {
	return &UserApi{
		userController: userController,
	}
}

func (api *UserApi) Register(ctx *gin.Context) {
	token, err := api.userController.Register(ctx)
	if err != nil {
		ctx.JSON(http.StatusConflict, &schemas.Response{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, &schemas.JWT{
		Token: token,
	})
}