package controllers

import (
	"github.com/gin-gonic/gin"

	"area51/services"
)

type UserController interface {
	Login(ctx *gin.Context) (string, error)
	Register(ctx *gin.Context) (string, error)
}

type userController struct {
	userService services.UserService
}