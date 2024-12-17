package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/services"
)

type ActionController interface {
	CreateAction(ctx *gin.Context) (string, error)
}

type actionController struct {
	service services.ActionService
}

func NewActionController(service services.ActionService) ActionController {
	return &actionController{
		service: service,
	}
}

func (controller *actionController) CreateAction(ctx *gin.Context) (string, error) {
	return controller.service.CreateAction(ctx)
}
