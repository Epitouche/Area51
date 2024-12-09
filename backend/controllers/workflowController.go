package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
)

type WorkflowController interface {
	CreateWorkflow(ctx *gin.Context) (string, error)
	GetMostRecentReaction(ctx *gin.Context) ([]schemas.GithubListCommentsResponse, error)
	AboutJson(ctx *gin.Context) (allWorkflows []schemas.WorkflowJson, err error)
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

func (controller *workflowController) AboutJson(ctx *gin.Context) (allWorkflowsJson []schemas.WorkflowJson, err error) {
	allWorkflows := controller.service.FindAll()
	for _, oneWorkflow := range allWorkflows {
		allWorkflowsJson = append(allWorkflowsJson, schemas.WorkflowJson{
			Name: oneWorkflow.Name,
			ActionId: oneWorkflow.ActionId,
			ReactionId: oneWorkflow.ReactionId,
			IsActive: oneWorkflow.IsActive,
			CreatedAt: oneWorkflow.CreatedAt,
		})
	}
	return allWorkflowsJson, nil
}

func (controller *workflowController) GetMostRecentReaction(ctx *gin.Context) ([]schemas.GithubListCommentsResponse, error) {
	return controller.service.GetMostRecentReaction(ctx)
}
