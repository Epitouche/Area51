package services

import (
	"area51/repository"
	"area51/schemas"
)

type ServicesService interface {
	FindAll() (allService []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	FindById(serviceId uint64) schemas.Service
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken)
	GetServices() []interface{}
	GetAllServices() (allServicesJson []schemas.ServiceJson, err error)
}

type ServiceInterface interface {
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken)
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
) ServicesService {
	newService := servicesService{
		repository: repository,
		allServicesSchema: []schemas.Service{
			{
				Name:        schemas.Github,
				Description: "This is the Github service",
				Image:       "https://pngimg.com/uploads/github/github_PNG80.png",
			},
			{
				Name:        schemas.Spotify,
				Description: "This is the Spotify Service",
				Image:       "tmp",
			},
		},
		allServices: []interface{}{
			githubService,
			spotifyService,
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

func (service *servicesService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64) {
	for _, oneService := range service.allServices {
		if oneService.(ServiceInterface).FindActionByName(name) != nil {
			return oneService.(ServiceInterface).FindActionByName(name)
		}
	}
	return nil
}

func (service *servicesService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken) {
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
