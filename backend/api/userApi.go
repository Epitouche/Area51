package api

import "area51/controllers"

type Dependencies struct {
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