package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
)

type ServicesController interface {
	AboutJson(ctx *gin.Context) (allServices []schemas.ServiceJson, err error)
}

type servicesController struct {
	service         services.ServicesService
	serviceAction   services.ActionService
	serviceReaction services.ReactionService
}

func NewServiceController(
	service services.ServicesService,
	serviceAction services.ActionService,
	serviceReaction services.ReactionService,
) ServicesController {
	return &servicesController{
		service:         service,
		serviceAction:   serviceAction,
		serviceReaction: serviceReaction,
	}
}

func (controller *servicesController) AboutJson(*gin.Context) (allServicesJson []schemas.ServiceJson, err error) {
	allServices := controller.service.FindAll()
	for _, oneService := range allServices {
		allServicesJson = append(allServicesJson, schemas.ServiceJson{
			Name:        schemas.ServiceName(oneService.Name),
			Description: oneService.Description,
			Action:      controller.serviceAction.GetAllServicesByServiceId(oneService.Id),
			Reaction:    controller.serviceReaction.GetAllServicesByServiceId(oneService.Id),
			Image:       oneService.Image,
			IsOAuth:     oneService.IsOAuth,
		})
	}
	return allServicesJson, nil
}
