package controllers

import "area51/services"

type ActionController interface{}

type actionController struct {
	service services.ActionService
}

func NewActionController(service services.ActionService) ActionController {
	return &actionController{
		service: service,
	}
}
