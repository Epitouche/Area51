package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
	GetAccessToken(ctx *gin.Context) (string, error)
}

type userController struct {
	userService services.UserService
	jWtService  services.JWTService
}

func NewUserController(userService services.UserService,
	jWtService services.JWTService,
) UserController {
	return &userController{
		userService: userService,
		jWtService:  jWtService,
	}
}

func (controller *userController) Login(ctx *gin.Context) (string, error) {
	var credentials schemas.LoginCredentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", err
	}

	newUser := schemas.User{
		Username: credentials.Username,
		Password: credentials.Password,
	}

	token, err := controller.userService.Login(newUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (controller *userController) Register(ctx *gin.Context) (string, error) {
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
		Email:    credentials.Email,
		Password: credentials.Password,
	}
	token, err := controller.userService.Register(newUser)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (controller *userController) GetAccessToken(ctx *gin.Context) (string, error) {
	// var credentials schemas.MobileToken
	cookies, _ := ctx.Request.Cookie("token")
	token := cookies.Name
	fmt.Printf("token: %v\n", token)
	// credentials.Token = token
	return token, nil
}
