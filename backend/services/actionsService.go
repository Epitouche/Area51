package services

import (
	"fmt"

	"area51/repository"
	"area51/schemas"
)

type ActionService interface {
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
	serviceService ServicesService
}

func NewActionService(
	repository repository.ActionRepository,
	serviceService ServicesService,
) ActionService {
	newActionService := &actionService{
		repository:     repository,
		serviceService: serviceService,
	}
	newActionService.SaveAllAction()
	return newActionService
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