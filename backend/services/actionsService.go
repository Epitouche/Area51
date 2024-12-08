package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"area51/repository"
	"area51/schemas"
)

type ActionService interface {
	CreateAction(ctx *gin.Context) (string, error)
	FindAll() []schemas.Action
	SaveAllAction()
	FindById(actionId uint64) schemas.Action
	GetAllServicesByServiceId(serviceId uint64) (actionJson []schemas.ActionJson)
}

type ServiceAction interface {
	GetServiceActionInfo() []schemas.Action
}

type actionService struct {
	repository     repository.ActionRepository
	userService    UserService
	serviceService ServicesService
}

func NewActionService(
	repository repository.ActionRepository,
	serviceService ServicesService,
	userService UserService,
) ActionService {
	newActionService := &actionService{
		repository:     repository,
		serviceService: serviceService,
		userService: userService,
	}
	newActionService.SaveAllAction()
	return newActionService
}

func (service *actionService) CreateAction(ctx *gin.Context) (string, error) {
	var result schemas.ActionResult
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	_, err = service.userService.GetUserInfos(tokenString)
	if err != nil {
		return "", err
	}
	serviceInfo := service.serviceService.FindByName(schemas.Github)

	newAction := schemas.Action{
		Name: result.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Description: result.Description,
		ServiceId: serviceInfo.Id,
	}
	service.repository.Save(newAction)
	return "Action Created successfully", nil
}

func (service *actionService) FindAll() []schemas.Action {
	return service.repository.FindAll()
}

func (service *actionService) GetAllServicesByServiceId(
	serviceId uint64,
) (actionJson []schemas.ActionJson) {
	allActionForService := service.repository.FindByServiceId(serviceId)
	for _, oneAction := range allActionForService {
		actionJson = append(actionJson, schemas.ActionJson{
			Name:        oneAction.Name,
			Description: oneAction.Description,
		})
	}
	return actionJson
}

func (service *actionService) SaveAllAction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceAction, ok := services.(ServiceAction); ok {
			actions := serviceAction.GetServiceActionInfo()
			for _, action := range actions {
				service.repository.Save(action)
			}
		} else {
			fmt.Println("Service is not ServiceAction")
		}
	}
}

func (service *actionService) FindById(actionId uint64) schemas.Action {
	return service.repository.FindById(actionId)
}
