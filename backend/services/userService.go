package services

import (
	"errors"
	"fmt"
	"strconv"

	"area51/database"
	"area51/repository"
	"area51/schemas"
)

type UserService interface {
	Login(user schemas.User) (JWTtoken string, err error)
	Register(newUser schemas.User) (JWTtoken string, err error)
	GetUserInfos(token string) (userInfos schemas.User, err error)
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

func (service *userService) Login(newUser schemas.User) (JWTtoken string, err error) {
	user := service.repository.FindByUsername(newUser.Username)
	if user.Username == "" {
		return "", errors.New("User not found")
	}
	if database.CompareHashAndPassword(user.Password, newUser.Password) {
		return service.serviceJWT.GenerateJWTToken(
			strconv.FormatUint(user.Id, 10),
			user.Username,
			user.IsAdmin,
		), nil
	}


	return "", errors.New("Invalid password")

}

func (service *userService) Register(newUser schemas.User) (JWTtoken string, err error) {
	user := service.repository.FindByEmail(newUser.Email)
	if user.Email != "" {
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
	return service.serviceJWT.GenerateJWTToken(fmt.Sprint(newUser.Id), newUser.Username, false), nil
}

func (service *userService) GetUserInfos(token string) (userInfos schemas.User, err error) {
	userId, err := service.serviceJWT.GetUserIdFromToken(token)
	if err != nil {
		return schemas.User{}, err
	}
	userInfos = service.repository.FindByID(userId)
	return userInfos, nil
}