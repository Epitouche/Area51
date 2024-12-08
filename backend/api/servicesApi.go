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
	serviceController controllers.ServicesController
	workflowController controllers.WorkflowController
}

func NewServicesApi(serviceController controllers.ServicesController, workflowController controllers.WorkflowController) *ServicesApi {
	return &ServicesApi{
		serviceController: serviceController,
		workflowController: workflowController,
	}
}

func (api *ServicesApi) AboutJson(ctx *gin.Context) {
	allServices, err := api.serviceController.AboutJson(ctx)
	allWorkflows, err := api.workflowController.AboutJson(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &schemas.BasicResponse{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"client": map[string]string{
				"host": ctx.ClientIP(),
			},
			"server": map[string]any{
				"current_time": fmt.Sprintf("%d", time.Now().Unix()),
				"services":     allServices,
				"workflows":	allWorkflows,
			},
		})
	}
}