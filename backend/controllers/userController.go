package controllers

import (
	"area51/schemas"
	"area51/services"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context) (string, error)
}

type userControllerStruct struct {
	userService services.UserService
	jwtService services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userControllerStruct{
		userService: userService,
		jwtService: jwtService,
	}
}

func (controller *userControllerStruct) Register(ctx *gin.Context) (string, error) {
	var credentials schemas.RegisterCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", err
	}
	if len(credentials.Username) < 4 {
		return "", errors.New("username must be at least 4 characters long")
	}
	if len(credentials.Password) < 8 {
		return "", errors.New("password must be at least 8 characters long")
	}
	if len(credentials.Email) < 4 {
		return "", errors.New("email must be at least 4 characters long")
	}

	newUser := schemas.User{
		Username: credentials.Username,
		Email: credentials.Email,
		Password: credentials.Password,
	}
	token, err := controller.userService.Register(newUser)
	if err != nil {
		return "", err
	}
	return token, nil
}