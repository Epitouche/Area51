package api

import (
	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
	"area51/toolbox"
)

type WorkflowApi struct {
	workflowController controllers.WorkflowController
}

func NewWorkflowApi(controller controllers.WorkflowController) *WorkflowApi {
	return &WorkflowApi{
		workflowController: controller,
	}
}

func (api *WorkflowApi) CreateWorkflow(ctx *gin.Context) {
	token, err := api.workflowController.CreateWorkflow(ctx)
	toolbox.HandleError(ctx, err, schemas.BasicResponse{Message: token})
}

func (api *WorkflowApi) ActivateWorkflow(ctx *gin.Context) {
	err := api.workflowController.ActivateWorkflow(ctx)
	toolbox.HandleError(ctx, err, schemas.BasicResponse{Message: "Workflow Status Updated"})
}

func (api *WorkflowApi) GetMostRecentReaction(ctx *gin.Context) {
	reaction, err := api.workflowController.GetMostRecentReaction(ctx)
	toolbox.HandleError(ctx, err, reaction)
}

func (api *WorkflowApi) GetAllReactionsForAWorkflow(ctx *gin.Context) {
	reactions, err := api.workflowController.GetAllReactionsForAWorkflow(ctx)
	toolbox.HandleError(ctx, err, reactions)
}

func (api *WorkflowApi) DeleteWorkflow(ctx *gin.Context) {
	err := api.workflowController.DeleteWorkflow(ctx)
	toolbox.HandleError(ctx, err, schemas.BasicResponse{Message: "Workflow Deleted"})
}

func (api *WorkflowApi) UpdateWorkflow(ctx *gin.Context) {
	err := api.workflowController.UpdateWorkflow(ctx)
	toolbox.HandleError(ctx, err, schemas.BasicResponse{Message: "Workflow Updated"})
}
