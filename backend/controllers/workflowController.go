package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/services"
)

type WorkflowController interface {
	CreateWorkflow(ctx *gin.Context) (string, error)
}

type workflowController struct {
	service services.WorkflowService
}

func NewWorkflowController(service services.WorkflowService) WorkflowController {
	return &workflowController{
		service: service,
	}
}

func (controller *workflowController) CreateWorkflow(ctx *gin.Context) (string, error) {
	return controller.service.CreateWorkflow(ctx)
}

