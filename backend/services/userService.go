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
	UpdateUserInfos(newUser schemas.User) error
	GetUserById(userId uint64) schemas.User
	GetUserByUsername(username string) schemas.User
	GetUserByEmail(email string) schemas.User
	CreateUser(newUser schemas.User) error
	GetAllServices(userId uint64) ([]schemas.ServiceToken, error)
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
		return "", errors.New("user not found")
	}

	if database.CompareHashAndPassword(user.Password, newUser.Password) {
		return service.serviceJWT.GenerateJWTToken(
			strconv.FormatUint(user.Id, 10),
			user.Username,
			user.IsAdmin,
		), nil
	}

	if user.Username == newUser.Username {
		if newUser.TokenId != nil && *newUser.TokenId != 0 {
			return service.serviceJWT.GenerateJWTToken(
				strconv.FormatUint(user.Id, 10),
				user.Username,
				false,
			), nil
		}
	}

	return "", errors.New("invalid password")
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

	userInfos = service.repository.FindById(userId)
	return userInfos, nil
}

func (service *userService) UpdateUserInfos(newUser schemas.User) error {
	service.repository.Update(newUser)
	return nil
}

func (service *userService) GetUserById(userId uint64) schemas.User {
	return service.repository.FindById(userId)
}

func (service *userService) GetUserByUsername(username string) schemas.User {
	return service.repository.FindByUsername(username)
}

func (service *userService) GetUserByEmail(email string) schemas.User {
	return service.repository.FindByEmail(email)
}

func (service *userService) CreateUser(newUser schemas.User) error {
	service.repository.Save(newUser)
	return nil
}

func (service *userService) GetAllServices(userId uint64) ([]schemas.ServiceToken, error) {
	return service.repository.FindAllServicesByUserId(userId), nil
}
