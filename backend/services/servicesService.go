package services

import (
	"area51/repository"
	"area51/schemas"
)

type ServicesService interface {
	FindAll() (allService []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64)
	FindReactionByName(name string) func (workflowId uint64)
	GetServices() []interface{}
	GetAllServices() (allServicesJson []schemas.ServiceJson, err error)
}

type servicesService struct {
	repository 			repository.ServiceRepository
	allServices 		[]interface{}
	allServicesSchema 	[]schemas.Service
}

func NewServicesService(
	repository repository.ServiceRepository,
	githubService GithubService,
	) ServicesService {
	newService := servicesService{
		repository: repository,
		allServicesSchema: []schemas.Service{
			{
				Name: schemas.Github,
				Description: "This is a code storage service",
			},
		},
		allServices: []interface{}{githubService},
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
		if githubService, ok := oneService.(GithubService); ok {
			return githubService.FindActionByName(name)
		}
	}
	return nil
}

func (service *servicesService) FindReactionByName(name string) func (workflowId uint64) {
	for _, oneService := range service.allServices {
		if githubService, ok := oneService.(GithubService); ok {
			return githubService.FindReactionByName(name)
		}
	}
	return nil
}

