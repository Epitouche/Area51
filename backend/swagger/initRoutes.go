package swagger

import (
	"net/http"

	"area51/api"
	"area51/schemas"
)

type Dependencies struct {
	UserApi *api.UserApi
	GithubApi *api.GithubApi
	ServicesApi *api.ServicesApi
	WorkflowApi *api.WorkflowApi
	ActionApi *api.ActionApi
	SpotifyApi *api.SpotifyApi
	MobileApi *api.MobileApi
	MicrosoftApi *api.MicrosoftApi
	GoogleApi *api.GoogleApi
}

type SwaggerRoutes interface {
	InitRoutes(deps Dependencies)
}

func InitRoutes(deps Dependencies) {
	var routes = []schemas.Route{
		{
			Path: "/about.json",
			Method: "GET",
			Handler: deps.ServicesApi.AboutJson,
			Description: "Get infos of the App",
			Product: []string{"application/json"},
			Tags: []string{"about.json"},
			ParamQueryType: "formData",
			Params: map[string]string{
			},
			Responses: map[int][]string{
				http.StatusOK: {
					"test",
					"test",
				},
			},
		},
	}
	ImpactSwaggerFiles(routes)
}