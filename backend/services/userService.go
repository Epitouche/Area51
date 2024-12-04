package services

import (
	"area51/database"
	"area51/repository"
	"area51/schemas"
	"errors"
	"fmt"
)

type UserService interface {
	Register(newUser schemas.User) (JWTtoken string, err error)
}

type userServiceStruct struct {
	authorizedUsername string
	authorizedPassword string
	repository repository.UserRepository
	serviceJWT JWTService
}

func NewUserService(userRepository repository.UserRepository, serviceJWT JWTService) UserService {
	return &userServiceStruct{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository: userRepository,
		serviceJWT: serviceJWT,
	}
}

func (service *userServiceStruct) Register(newUser schemas.User) (JWTtoken string, err error) {
	user := service.repository.FindByEmail(newUser.Email)
	if len(user) != 0 {
		return "", errors.New("email already in use")
	}

	if newUser.Password != "" {
		hashedPassword, err := database.HashPassword(newUser.Password)
		if err != nil {
			return "", errors.New("error while hashing the password")
		}
		newUser.Password = hashedPassword
	}
	service.repository.Save(newUser)
	return service.serviceJWT.GenerateToken(fmt.Sprint(newUser.Id), newUser.Username, false), nil
}