package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, schemas.BasicResponse{Message: token})
}

func (api *WorkflowApi) GetMostRecentReaction(ctx *gin.Context) {
	reaction, err := api.workflowController.GetMostRecentReaction(ctx)
	if err != nil && err.Error() == "no authorization header found" {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, reaction)
}
