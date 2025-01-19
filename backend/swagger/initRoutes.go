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
			Path:        "/about.json",
			Method:      "GET",
			Handler:     deps.ServicesApi.AboutJson,
			Description: "Get the AREA51 services",
			Product:     []string{"application/json"},
			Tags:        []string{"about.js"},
			ParamQueryType: "formData",
			Params: map[string]string{
			},
			Responses: map[int][]string{
				http.StatusOK: {
					"About",
					"schemas.AboutJSON",
				},
				http.StatusInternalServerError: {
					"Internal Server Error",
					"schemas.BasicResponse",
				},
			},
        },
	}

	filePath := "docs/swagger.json"

	pathsOfRoutesWanted := []string{}
	existingRoutes := ExtractExistingRoutes(filePath)
	for _, route := range routes {
		pathsOfRoutesWanted = append(pathsOfRoutesWanted, route.Path)
	}
	routesToRemove := FindRoutesToRemove(existingRoutes, pathsOfRoutesWanted)
	if len(routesToRemove) > 0 || (len(pathsOfRoutesWanted) != len(existingRoutes)) {
		ImpactSwaggerFiles(routes, routesToRemove)
	}
}