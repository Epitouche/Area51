package services

import (
	"fmt"

	"area51/repository"
	"area51/schemas"
)

type ServicesService interface {
	FindAll() (allService []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	FindById(serviceId uint64) schemas.Service
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
	GetServices() []interface{}
	GetAllServices() (allServicesJson []schemas.ServiceJson, err error)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
}

type ServiceInterface interface {
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
}

type servicesService struct {
	repository        repository.ServiceRepository
	allServices       []interface{}
	allServicesSchema []schemas.Service
}

func NewServicesService(
	repository repository.ServiceRepository,
	githubService GithubService,
	spotifyService SpotifyService,
	googleService GoogleService,
	microsoftService MicrosoftService,
	weatherService WeatherService,
	interpolService InterpolService,
) ServicesService {
	newService := servicesService{
		repository: repository,
		allServicesSchema: []schemas.Service{
			{
				Name:        schemas.Github,
				Description: "This is the Github service",
				Image:       "https://pngimg.com/uploads/github/github_PNG80.png",
				IsOAuth:     true,
			},
			{
				Name:        schemas.Spotify,
				Description: "This is the Spotify Service",
				Image:       "https://www.freepnglogos.com/uploads/spotify-logo-png/spotify-logo-spotify-symbol-3.png",
				IsOAuth:     true,
			},
			{
				Name:        schemas.Google,
				Description: "This is the Google service",
				Image:       "https://pngimg.com/uploads/google/google_PNG19630.png",
				IsOAuth:     true,
			},
			{
				Name:        schemas.Microsoft,
				Description: "This is the Microsoft Service",
				Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/4/44/Microsoft_logo.svg/1024px-Microsoft_logo.svg.png",
				IsOAuth:     true,
			},
			{
				Name:        schemas.Weather,
				Description: "This is the Weather Service",
				Image:       "https://img.icons8.com/?size=100&id=15359&format=png&color=000000",
			},
			{
				Name:        schemas.Interpol,
				Description: "This is the Interpol Service",
				Image:       "tmp",
			},
		},
		allServices: []interface{}{
			githubService,
			spotifyService,
			googleService,
			microsoftService,
			weatherService,
			interpolService,
		},
	}
	newService.InitialSaveService()
	return &newService
}

func (service *servicesService) InitialSaveService() {
	for _, oneService := range service.allServicesSchema {
		serviceByName := service.repository.FindAllByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
		}
	}
}

func (service *servicesService) FindAll() (allService []schemas.Service) {
	return service.repository.FindAll()
}

func (service *servicesService) FindByName(serviceName schemas.ServiceName) schemas.Service {
	return service.repository.FindByName(serviceName)
}

func (service *servicesService) GetServices() []interface{} {
	return service.allServices
}

func (service *servicesService) GetAllServices() (allServicesJson []schemas.ServiceJson, err error) {
	allServices := service.repository.FindAll()

	for _, oneService := range allServices {
		allServicesJson = append(allServicesJson, schemas.ServiceJson{
			Name: schemas.ServiceName(oneService.Name),
		})
	}
	return allServicesJson, nil
}

func (service *servicesService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	for _, oneService := range service.allServices {
		fmt.Println(oneService)
		if oneService.(ServiceInterface).FindActionByName(name) != nil {
			return oneService.(ServiceInterface).FindActionByName(name)
		}
	}
	return nil
}

func (service *servicesService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	for _, oneService := range service.allServices {
		if oneService.(ServiceInterface).FindReactionByName(name) != nil {
			return oneService.(ServiceInterface).FindReactionByName(name)
		}
	}
	return nil
}

func (service *servicesService) FindById(serviceId uint64) schemas.Service {
	return service.repository.FindById(serviceId)
}

func (service *servicesService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	for i, s := range service.allServicesSchema {
		for j, oneService := range service.allServices {
			if s.Name == serviceName {
				if j == i {
					if oneService.(ServiceInterface).GetUserInfosByToken(accessToken, serviceName) != nil {
						return oneService.(ServiceInterface).GetUserInfosByToken(accessToken, serviceName)
					}
				}
			}
		}
	}
	return nil
}
