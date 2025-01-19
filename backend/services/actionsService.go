package services

import (
	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type ActionService interface {
	FindAll() []schemas.Action
	SaveAllAction()
	FindById(actionId uint64) schemas.Action
	GetAllServicesByServiceId(serviceId uint64) (actionJson []schemas.ActionJson)
}

type ServiceAction interface {
	GetServiceActionInfo() []schemas.Action
}

type actionService struct {
	repository       repository.ActionRepository
	userService      UserService
	serviceService   ServicesService
	allActions       []interface{}
	allActionsSchema []schemas.Action
}

func NewActionService(
	repository repository.ActionRepository,
	serviceService ServicesService,
	userService UserService,
) ActionService {
	newActionService := &actionService{
		repository:     repository,
		serviceService: serviceService,
		userService:    userService,
		allActionsSchema: []schemas.Action{
			{
				Name:        string(schemas.GithubPullRequest),
				Description: "Creation or deletion of a pull request",
				ServiceId:   serviceService.FindByName(schemas.Github).Id,
				Options: toolbox.RealObject(schemas.GithubPullRequestOptions{
					Owner: "my github username",
					Repo:  "name of the repository",
				}),
			},
			{
				Name:        string(schemas.GithubPushOnRepo),
				Description: "Detect a push on a repository",
				ServiceId:   serviceService.FindByName(schemas.Github).Id,
				Options: toolbox.RealObject(schemas.GithubPushOnRepoOptions{
					Owner:  "my github username",
					Repo:   "name of the repository",
					Branch: "main",
				}),
			},
			{
				Name:        string(schemas.SpotifyAddTrackAction),
				Description: "Add a track to a playlist",
				ServiceId:   serviceService.FindByName(schemas.Spotify).Id,
				Options: toolbox.RealObject(schemas.SpotifyActionOptionsInfo{
					PlaylistURL: "https://open.spotify.com/playlist/37i9dQZF1DXcBWIGoYBM5M",
				}),
			},
			{
				Name:        string(schemas.GoogleGetEmailAction),
				Description: "Get the email of the user",
				ServiceId:   serviceService.FindByName(schemas.Google).Id,
				Options: toolbox.RealObject(schemas.GoogleActionOptions{
					Label: "name of the box (INBOX, SPAM, ...)",
				}),
			},
			{
				Name:        string(schemas.MicrosoftOutlookEventsAction),
				Description: "Detect an event in the oulook calendar of the user",
				ServiceId:   serviceService.FindByName(schemas.Microsoft).Id,
				Options: toolbox.RealObject(schemas.MicrosoftOutlookEventsOptions{
					Subject: "Réunion de travail",
				}),
			},
			{
				Name:        string(schemas.WeatherCurrentAction),
				Description: "Get the current weather",
				ServiceId:   serviceService.FindByName(schemas.Weather).Id,
				Options: toolbox.RealObject(schemas.WeatherCurrentOptions{
					CityName:     "Bordeaux",
					LanguageCode: "FR",
					Temperature:  "0",
					CompareSign:  "> or < or =",
				}),
			},
			{
				Name:        string(schemas.WeatherTimeAction),
				Description: "Wait for a specific time",
				ServiceId:   serviceService.FindByName(schemas.Weather).Id,
				Options: toolbox.RealObject(schemas.WeatherSpecificTimeOption{
					DateTime: "2025-01-18",
					CityName: "Bordeaux",
				}),
			},
			{
				Name:        string(schemas.MicrosoftTeamGroup),
				Description: "Modify a Teams group",
				ServiceId:   serviceService.FindByName(schemas.Microsoft).Id,
				Options: toolbox.RealObject(schemas.MicrosoftTeamsGroupOptionsInfos{
					Name: "Area51",
				}),
			},
			{
				Name:        string(schemas.InterpolNewRedNotice),
				Description: "Verify if the number of red notices has changed",
				ServiceId:   serviceService.FindByName(schemas.Interpol).Id,
				Options: toolbox.RealObject(schemas.InterpolActionOptions{
					SexId: "M or F or U",
				}),
			},
		},
		allActions: []interface{}{serviceService},
	}
	newActionService.SaveAllAction()
	return newActionService
}

func (service *actionService) FindAll() []schemas.Action {
	return service.repository.FindAll()
}

func (service *actionService) GetAllServicesByServiceId(
	serviceId uint64,
) (actionJson []schemas.ActionJson) {
	allActionForService := service.repository.FindByServiceId(serviceId)

	for _, oneAction := range allActionForService {
		actionJson = append(actionJson, schemas.ActionJson{
			Name:        oneAction.Name,
			Description: oneAction.Description,
			Options:     oneAction.Options,
			ActionId:    oneAction.Id,
		})
	}
	return actionJson
}

func (service *actionService) SaveAllAction() {
	for _, oneService := range service.allActionsSchema {
		serviceByName := service.repository.FindAllByName(oneService.Name)
		if len(serviceByName) == 0 {
			service.repository.Save(oneService)
		}
	}
}

func (service *actionService) FindById(actionId uint64) schemas.Action {
	return service.repository.FindById(actionId)
}
