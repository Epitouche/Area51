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
	Login(user schemas.User, actualService schemas.Service) (JWTtoken string, err error)
	Register(newUser schemas.User) (JWTtoken string, err error)
	GetUserInfos(token string) (userInfos schemas.User, err error)
	UpdateUserInfos(newUser schemas.User) error
	GetUserById(userId uint64) schemas.User
	GetUserByUsername(username string) schemas.User
	GetUserByEmail(email *string) schemas.User
	CreateUser(newUser schemas.User) error
	DeleteUser(userId uint64) error
	AddServiceToUser(user schemas.User, serviceToAdd schemas.ServiceToken) error
	GetAllServicesForUser(userId uint64) ([]schemas.ServiceToken, error)
	GetServiceByIdForUser(user schemas.User, serviceId uint64) (schemas.ServiceToken, error)
	GetAllServices(userId uint64) ([]schemas.ServiceToken, error)
	GetAllWorkflows(userId uint64) ([]schemas.Workflow, error)
	LogoutFromService(userId uint64, serviceToDelete schemas.Service) error
}

type userService struct {
	authorizedUsername string
	authorizedPassword string
	repository         repository.UserRepository
	serviceJWT         JWTService
}

func NewUserService(
	repository repository.UserRepository,
	serviceJWT JWTService,
) UserService {
	return &userService{
		authorizedUsername: "root",
		authorizedPassword: "password",
		repository:         repository,
		serviceJWT:         serviceJWT,
	}
}

func (service *userService) Login(newUser schemas.User, actualService schemas.Service) (JWTtoken string, err error) {
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
		serviceToken, _ := service.GetServiceByIdForUser(user, actualService.Id)
		if serviceToken.Id != 0 {
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
	if user.Email != nil {
		return "", errors.New("email already in use")
	}

	if newUser.Password != nil {
		hashedPassword, err := database.HashPassword(*newUser.Password)
		if err != nil {
			return "", errors.New("error while hashing the password")
		}
		newUser.Password = &hashedPassword
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

func (service *userService) GetUserByEmail(email *string) schemas.User {
	return service.repository.FindByEmail(email)
}

func (service *userService) CreateUser(newUser schemas.User) error {
	service.repository.Save(newUser)
	return nil
}

func (service *userService) GetAllServices(userId uint64) ([]schemas.ServiceToken, error) {
	return service.repository.FindAllServicesByUserId(userId), nil
}

func (service *userService) GetAllWorkflows(userId uint64) ([]schemas.Workflow, error) {
	return service.repository.FindAllWorkflowsByUserId(userId), nil
}

func (service *userService) AddServiceToUser(user schemas.User, serviceToAdd schemas.ServiceToken) error {
	service.repository.AddServiceToUser(user, serviceToAdd)
	return nil
}

func (service *userService) GetAllServicesForUser(userId uint64) ([]schemas.ServiceToken, error) {
	return service.repository.GetAllServicesForUser(userId)
}

func (service *userService) GetServiceByIdForUser(user schemas.User, serviceId uint64) (schemas.ServiceToken, error) {
	return service.repository.GetServiceByIdForUser(user, serviceId)
}

func (service *userService) LogoutFromService(userId uint64, serviceToDelete schemas.Service) error {
	user := service.GetUserById(userId)
	return service.repository.LogoutFromService(user, serviceToDelete)
}

func (service *userService) DeleteUser(userId uint64) error {
	user := service.GetUserById(userId)
	service.repository.Delete(user)
	return nil
}
