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
	if err != nil && err.Error() == "no authorization header found" {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, schemas.BasicResponse{Message: token})
}

func (api *WorkflowApi) ActivateWorkflow(ctx *gin.Context) {
	err := api.workflowController.ActivateWorkflow(ctx)
	if err != nil && err.Error() == "no authorization header found" {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, schemas.BasicResponse{Message: "Workflow State Updated"})
}

func (api *WorkflowApi) GetMostRecentReaction(ctx *gin.Context) {
	reaction, err := api.workflowController.GetMostRecentReaction(ctx)
	if err != nil && err.Error() == "no authorization header found" {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, schemas.BasicResponse{
			Message: err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, reaction)
	}
}

func (api *WorkflowApi) GetAllReactionsForAWorkflow(ctx *gin.Context) {
	reactions, err := api.workflowController.GetAllReactionsForAWorkflow(ctx)
	if err != nil && err.Error() == "no authorization header found" {
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, reactions)
	}
}

func (api *WorkflowApi) DeleteWorkflow(ctx *gin.Context) {
	err := api.workflowController.DeleteWorkflow(ctx)
	if err != nil && err.Error() == "no authorization header found" {
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, schemas.BasicResponse{Message: "Workflow Deleted"})
}

func (api *WorkflowApi) UpdateWorkflow(ctx *gin.Context) {
	err := api.workflowController.UpdateWorkflow(ctx)
	if err != nil && err.Error() == "workflow not found" {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, schemas.BasicResponse{
			Message: "Workflow updated",
		})
	}
}