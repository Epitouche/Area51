package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"area51/services"
)

type WorkflowController interface {
	CreateWorkflow(ctx *gin.Context) (string, error)
	ActivateWorkflow(ctx *gin.Context) error
	GetMostRecentReaction(ctx *gin.Context) ([]json.RawMessage, error)
	GetAllReactionsForAWorkflow(ctx *gin.Context) ([]json.RawMessage, error)
	DeleteWorkflow(ctx *gin.Context) error
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

func (controller *workflowController) GetMostRecentReaction(ctx *gin.Context) ([]json.RawMessage, error) {
	return controller.service.GetMostRecentReaction(ctx)
}

func (controller *workflowController) DeleteWorkflow(ctx *gin.Context) error {
	return controller.service.DeleteWorkflow(ctx)
}

func (controller *workflowController) GetAllReactionsForAWorkflow(ctx *gin.Context) ([]json.RawMessage, error) {
	return controller.service.GetAllReactionsForAWorkflow(ctx)
}
