package services

import (
	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type ReactionService interface {
	FindAll() []schemas.Reaction
	SaveAllReaction()
	FindById(reactionId uint64) schemas.Reaction
	UpdateTrigger(reaction schemas.Reaction)
	GetAllServicesByServiceId(serviceId uint64) (reactionJson []schemas.ReactionJson)
}

type ServiceReaction interface {
	GetServiceReactionInfo() []schemas.Reaction
}

type reactionService struct {
	repository         repository.ReactionRepository
	serviceService     ServicesService
	allReactions       []interface{}
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
				Name:        "list_comments",
				Description: "List all comments of a repository",
				ServiceId:   serviceService.FindByName(schemas.Github).Id,
			},
			{
				Name:        "add_track_reaction",
				Description: "Add a track to a playlist",
				ServiceId:   serviceService.FindByName(schemas.Spotify).Id,
				Options: toolbox.MustMarshal(schemas.SpotifyReactionOptions{
					PlaylistURL: "string",
					TrackURL: "string",
				}),
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
			Options: oneReaction.Options,
		})
	}
	return reactionJson
}

func (service *reactionService) SaveAllReaction() {
	for _, oneService := range service.allReactionsSchema {
		serviceByName := service.repository.FindAllByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
		}
	}
}

func (service *reactionService) FindById(reactionId uint64) schemas.Reaction {
	return service.repository.FindById(reactionId)
}

func (service *reactionService) UpdateTrigger(reaction schemas.Reaction) {
	service.repository.UpdateTrigger(reaction)
}
