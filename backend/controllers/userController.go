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
	GetAllServices(ctx *gin.Context) ([]schemas.Service, error)
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

func (controller *userController) GetAllServices(ctx *gin.Context) ([]schemas.Service, error) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) < len("Bearer ") {
		return nil, fmt.Errorf("invalid token")
	}
	tokenString := authHeader[len("Bearer "):]
	var allServices []schemas.Service
	userId, err := controller.jWtService.GetUserIdFromToken(tokenString)
	if err != nil {
		return nil, err
	}
	services, err := controller.userService.GetAllServices(userId)
	if len(services) == 0 {
		return nil, fmt.Errorf("no services found")
	}
	if err != nil {
		return nil, err
	}
	for _, service := range services {
		allServices = append(allServices, controller.servicesService.FindById(service.ServiceId))

	}
	return allServices, nil
}
