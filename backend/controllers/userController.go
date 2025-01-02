package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
	GetAllServices(ctx *gin.Context) ([]schemas.Service, error)
	GetAccessToken(ctx *gin.Context) (string, error)
}

type userController struct {
	userService     services.UserService
	jWtService      services.JWTService
	servicesService services.ServicesService
}

func NewUserController(
	userService services.UserService,
	jWtService services.JWTService,
	servicesService services.ServicesService,
) UserController {
	return &userController{
		userService:     userService,
		jWtService:      jWtService,
		servicesService: servicesService,
	}
}

func (controller *userController) Login(ctx *gin.Context) (string, error) {
	var credentials schemas.LoginCredentials
	if err := ctx.ShouldBind(&credentials); err != nil {
		return "", err
	}

	token, err := controller.userService.Login(schemas.User{
		Username: credentials.Username,
		Password: credentials.Password,
	})
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

	token, err := controller.userService.Register(schemas.User{
		Username: credentials.Username,
		Email:    credentials.Email,
		Password: credentials.Password,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (controller *userController) GetAccessToken(ctx *gin.Context) (string, error) {
	cookies, err := ctx.Request.Cookie("token")
	if err != nil {
		return "", err
	}
	token := cookies.Name
	return token, nil
}

func (controller *userController) GetAllServices(ctx *gin.Context) ([]schemas.Service, error) {
	bearer, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		panic(err.Error())
	}
	userId, err := controller.jWtService.GetUserIdFromToken(bearer)
	if err != nil {
		return nil, err
	}
	services, err := controller.userService.GetAllServices(userId)
	if err != nil {
		return nil, err
	}
	var allServices []schemas.Service
	for _, service := range services {
		allServices = append(allServices, controller.servicesService.FindById(service.ServiceId))
	}
	return allServices, nil
}
