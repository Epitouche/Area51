package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message" : "Pong",
		})
	})

	server.Run(":8080")
}
