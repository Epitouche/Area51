package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
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
	if token, err := api.workflowController.CreateWorkflow(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	} else {
		ctx.JSON(http.StatusOK, token)
	}
}

func (api *WorkflowApi) GetMostRecentReaction(ctx *gin.Context) {
	if reaction, err := api.workflowController.GetMostRecentReaction(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	} else {
		ctx.JSON(http.StatusOK, reaction)
	}
}
