package services

import (
	"area51/repository"
	"area51/schemas"
)

type ServicesService interface {
	FindAll() (allService []schemas.Service)
	FindByName(serviceName schemas.ServiceName) schemas.Service
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