package services

import (
	"errors"

	"github.com/gin-gonic/gin"

	"area51/database"
	"area51/repository"
	"area51/schemas"
)

type UserService interface {
	Login(user schemas.User) (schemas.JWT, error)
	Register(ctx *gin.Context) (string, error)
}

type userService struct {
	authorizedUsername string
	authorizedPassword string
	repository         repository.UserRepository
	serviceJWT         JWTService
}


func NewUserService(repository repository.UserRepository, serviceJWT JWTService) UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository:         repository,
		serviceJWT:         serviceJWT,
	}
}

func (service *userService) Login(newUser schemas.User) (schemas.JWT, error) {
	user := service.repository.FindByEmail(newUser.Email)
	if user.Email == "" {
		return schemas.JWT{}, errors.New("User not found")
	}

	if database.CompareHashAndPassword(user.Password, newUser.Password) {
		// return service.serviceJWT.GenerateJWTToken(fmt.Sprint(user.Id), user.Email, user.IsAdmin), nil
	}

	return schemas.JWT{}, errors.New("Invalid password")

}

func (service *userService) Register(ctx *gin.Context) (string, error) {
	return "Register", nil
}