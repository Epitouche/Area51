package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/schemas"
	"area51/services"
)

type WorkflowController interface {
	CreateWorkflow(ctx *gin.Context) (string, error)
	ActivateWorkflow(ctx *gin.Context) error
	GetMostRecentReaction(ctx *gin.Context) ([]schemas.GithubListCommentsResponse, error)
	AboutJson(ctx *gin.Context) (allWorkflows []schemas.WorkflowJson, err error)
}

type workflowController struct {
	service         services.WorkflowService
	reactionService services.ReactionService
	actionService   services.ActionService
}

func NewWorkflowController(
	service services.WorkflowService,
	reactionService services.ReactionService,
	actionService services.ActionService,
) WorkflowController {
	return &workflowController{
		service:         service,
		reactionService: reactionService,
		actionService:   actionService,
	}
}

func (controller *workflowController) CreateWorkflow(ctx *gin.Context) (string, error) {
	return controller.service.CreateWorkflow(ctx)
}

func (controller *workflowController) ActivateWorkflow(ctx *gin.Context) error {
	return controller.service.ActivateWorkflow(ctx)
}

func (controller *workflowController) AboutJson(ctx *gin.Context) (allWorkflowsJson []schemas.WorkflowJson, err error) {
	allWorkflows := controller.service.FindAll()
	for _, oneWorkflow := range allWorkflows {
		action := controller.actionService.FindById(oneWorkflow.ActionId)
		reaction := controller.reactionService.FindById(oneWorkflow.ReactionId)
		allWorkflowsJson = append(allWorkflowsJson, schemas.WorkflowJson{
			Name:         oneWorkflow.Name,
			WorkflowId:   oneWorkflow.Id,
			ActionName:   action.Name,
			ReactionName: reaction.Name,
			IsActive:     oneWorkflow.IsActive,
			CreatedAt:    oneWorkflow.CreatedAt,
		})
	}
	return allWorkflowsJson, nil
}

func (controller *workflowController) GetMostRecentReaction(ctx *gin.Context) ([]schemas.GithubListCommentsResponse, error) {
	return controller.service.GetMostRecentReaction(ctx)
}
