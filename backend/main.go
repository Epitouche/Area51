package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/about.json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "simple about.json route",
		})
	})

	return router
}

func main() {
	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}