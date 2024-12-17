package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"area51/controllers"
	"area51/schemas"
)

type ServicesApi struct {
	serviceController  controllers.ServicesController
	workflowController controllers.WorkflowController
}

func NewServicesApi(serviceController controllers.ServicesController, workflowController controllers.WorkflowController) *ServicesApi {
	return &ServicesApi{
		serviceController:  serviceController,
		workflowController: workflowController,
	}
}

func (api *ServicesApi) AboutJson(ctx *gin.Context) {
	allServices, serviceErr := api.serviceController.AboutJson(ctx)
	allWorkflows, workflowErr := api.workflowController.AboutJson(ctx)

	if serviceErr != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.BasicResponse{
			Message: serviceErr.Error(),
		})
	} else if workflowErr != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.BasicResponse{
			Message: workflowErr.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"client": map[string]string{
				"host": ctx.ClientIP(),
			},
			"server": map[string]any{
				"current_time": fmt.Sprintf("%d", time.Now().Unix()),
				"services":     allServices,
				"workflows":    allWorkflows,
			},
		})
	}
}
