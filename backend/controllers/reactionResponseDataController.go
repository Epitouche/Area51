package controllers

import "area51/services"

type ReactionResponseDataController interface {}

type reactionResponseDataController struct {
	service services.ReactionResponseDataService
}

func NewReactionResponseDataController(service services.ReactionResponseDataService) ReactionResponseDataController {
	return &reactionResponseDataController{
		service: service,
	}
}
