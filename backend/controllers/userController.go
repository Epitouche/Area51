package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
	"area51/toolbox"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
	GetAllServices(ctx *gin.Context) ([]schemas.Service, error)
	GetAllWorkflows(ctx *gin.Context) ([]schemas.WorkflowJson, error)
}

type userController struct {
	userService     services.UserService
	jWtService      services.JWTService
	servicesService services.ServicesService
	reactionService services.ReactionService
	actionService   services.ActionService
}

func NewUserController(
	userService services.UserService,
	jWtService services.JWTService,
	servicesService services.ServicesService,
	reactionService services.ReactionService,
	actionService services.ActionService,
) UserController {
	return &userController{
		userService:     userService,
		jWtService:      jWtService,
		servicesService: servicesService,
		reactionService: reactionService,
		actionService:   actionService,
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
	if len(services) == 0 {
		return nil, fmt.Errorf("no services found")
	}
	if err != nil {
		return nil, err
	}
	var allServices []schemas.Service
	for _, service := range services {
		allServices = append(allServices, controller.servicesService.FindById(service.ServiceId))
	}
	return allServices, nil
}

func (controller *userController) GetAllWorkflows(ctx *gin.Context) ([]schemas.WorkflowJson, error) {
	bearer, _ := toolbox.GetBearerToken(ctx)
	userId, err := controller.jWtService.GetUserIdFromToken(bearer)
	if err != nil {
		return nil, err
	}
	workflows, err := controller.userService.GetAllWorkflows(userId)
	if len(workflows) == 0 {
		return nil, fmt.Errorf("no services found")
	}
	if err != nil {
		return nil, err
	}
	var allWorkflows []schemas.WorkflowJson
	for _, workflow := range workflows {
		action := controller.actionService.FindById(workflow.ActionId)
		reaction := controller.reactionService.FindById(workflow.ReactionId)
		allWorkflows = append(allWorkflows, schemas.WorkflowJson{
			Name:         workflow.Name,
			WorkflowId:   workflow.Id,
			ActionId:     workflow.ActionId,
			ReactionId:   workflow.ReactionId,
			ActionName:   action.Name,
			ReactionName: reaction.Name,
			IsActive:     workflow.IsActive,
			CreatedAt:    workflow.CreatedAt,
		})

	}
	return allWorkflows, nil
}
