package services

import (
	"area51/repository"
	"area51/schemas"
)

type ReactionResponseDataService interface {
	Save(reactionResponseData schemas.ReactionResponseData)
	Update(reactionResponseData schemas.ReactionResponseData)
	Delete(reactionResponseData schemas.ReactionResponseData)
	FindAll() []schemas.ReactionResponseData
	FindByWorkflowId(workflowId uint64) []schemas.ReactionResponseData
}

type reactionResponseDataService struct {
	repository repository.ReactionResponseDataRepository
}

func NewReactionResponseDataService(
	repository repository.ReactionResponseDataRepository,
	) ReactionResponseDataService {
	return &reactionResponseDataService{
		repository: repository,
	}
}

func (service *reactionResponseDataService) Save(reactionResponseData schemas.ReactionResponseData) {
	service.repository.Save(reactionResponseData)
}

func (service *reactionResponseDataService) Update(reactionResponseData schemas.ReactionResponseData) {
	service.repository.Update(reactionResponseData)
}

func (service *reactionResponseDataService) Delete(reactionResponseData schemas.ReactionResponseData) {
	service.repository.Delete(reactionResponseData)
}

func (service *reactionResponseDataService) FindAll() []schemas.ReactionResponseData {
	return service.repository.FindAll()
}

func (service *reactionResponseDataService) FindByWorkflowId(workflowId uint64) []schemas.ReactionResponseData {
	return service.repository.FindByWorkflowId(workflowId)
}
