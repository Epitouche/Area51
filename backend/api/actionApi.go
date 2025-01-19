package api

import (
	"area51/controllers"
)

type ActionApi struct {
	actionController controllers.ActionController
}

func NewActionApi(controller controllers.ActionController) *ActionApi {
	return &ActionApi{
		actionController: controller,
	}
}
