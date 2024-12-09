package services

import (
	"area51/repository"
	"area51/schemas"
)

type ReactionService interface {
	FindAll() []schemas.Reaction
	SaveAllReaction()
	FindById(reactionId uint64) schemas.Reaction
	GetAllServicesByServiceId(serviceId uint64) (reactionJson []schemas.ReactionJson)
}

type ServiceReaction interface {
	GetServiceReactionInfo() []schemas.Reaction
}

type reactionService struct {
	repository     repository.ReactionRepository
	serviceService ServicesService
	allReactions   []interface{}
	allReactionsSchema []schemas.Reaction

}

func NewReactionService(
	repository repository.ReactionRepository,
	serviceService ServicesService,
) ReactionService {
	newService := &reactionService{
		repository:     repository,
		serviceService: serviceService,
		allReactionsSchema: []schemas.Reaction{
			{
				Name: "list_comments",
				Description: "List all comments of a repository",
				ServiceId: serviceService.FindByName(schemas.Github).Id,
			},
		},
		allReactions: []interface{}{serviceService},
	}
	newService.SaveAllReaction()
	return newService
}

func (service *reactionService) FindAll() []schemas.Reaction {
	return service.repository.FindAll()
}

func (service *reactionService) GetAllServicesByServiceId(
	serviceId uint64,
) (reactionJson []schemas.ReactionJson) {
	allRectionForService := service.repository.FindByServiceId(serviceId)
	for _, oneReaction := range allRectionForService {
		reactionJson = append(reactionJson, schemas.ReactionJson{
			Name:        oneReaction.Name,
			Description: oneReaction.Description,
			ReactionId:  oneReaction.Id,
		})
	}
	return reactionJson
}

func (service *reactionService) SaveAllReaction() {
	for _, services := range service.serviceService.GetServices() {
		if serviceReaction, ok := services.(ServiceReaction); ok {
			reactions := serviceReaction.GetServiceReactionInfo()
			for _, reaction := range reactions {
				service.repository.Save(reaction)
			}
		} else {
			println("ServiceReaction interface not implemented")
		}
	}
}

func (service *reactionService) FindById(reactionId uint64) schemas.Reaction {
	return service.repository.FindById(reactionId)
}