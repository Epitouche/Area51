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
	controller controllers.ServicesController
}

func NewServiceApi(controller controllers.ServicesController) *ServicesApi {
	return &ServicesApi{
		controller: controller,
	}
}

func (api *ServicesApi) AboutJson(ctx *gin.Context) {
	allServices, err := api.controller.AboutJson(ctx)

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
			},
		})
	}
}