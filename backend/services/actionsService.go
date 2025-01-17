package services

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type ActionService interface {
	CreateAction(ctx *gin.Context) (string, error)
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
				Options: toolbox.MustMarshal(schemas.GithubPullRequestOptions{
					Owner: "string",
					Repo:  "string",
				}),
			},
			{
				Name:        string(schemas.GithubPushOnRepo),
				Description: "Detect a push on a repository",
				ServiceId:   serviceService.FindByName(schemas.Github).Id,
				Options: toolbox.MustMarshal(schemas.GithubPushOnRepoOptions{
					Owner:  "string",
					Repo:   "string",
					Branch: "string",
				}),
			},
			{
				Name:        string(schemas.SpotifyAddTrackAction),
				Description: "Add a track to a playlist",
				ServiceId:   serviceService.FindByName(schemas.Spotify).Id,
				Options: toolbox.MustMarshal(schemas.SpotifyActionOptionsInfo{
					PlaylistURL: "string",
				}),
			},
			{
				Name:        string(schemas.GoogleGetEmailAction),
				Description: "Get the email of the user",
				ServiceId:   serviceService.FindByName(schemas.Google).Id,
				Options: toolbox.MustMarshal(schemas.GoogleActionOptions{
					Label: "string",
				}),
			},
			{
				Name:        string(schemas.MicrosoftOutlookEventsAction),
				Description: "Detect an event in the oulook calendar of the user",
				ServiceId:   serviceService.FindByName(schemas.Microsoft).Id,
				Options: toolbox.MustMarshal(schemas.MicrosoftOutlookEventsOptions{
					Subject: "string",
				}),
			},
			{
				Name:        string(schemas.WeatherCurrentAction),
				Description: "Get the current weather",
				ServiceId:   serviceService.FindByName(schemas.Weather).Id,
				Options: toolbox.MustMarshal(schemas.WeatherCurrentOptions{
					CityName:     "string",
					LanguageCode: "string",
					Temperature:  0,
					CompareSign:  "string",
				}),
			},
			{
				Name:        string(schemas.WeatherTimeAction),
				Description: "Wait for a specific time",
				ServiceId:   serviceService.FindByName(schemas.Weather).Id,
				Options: toolbox.MustMarshal(schemas.WeatherSpecificTimeOption{
					DateTime: "string",
					CityName: "string",
				}),
			},
		},
		allActions: []interface{}{serviceService},
	}
	newActionService.SaveAllAction()
	return newActionService
}

func (service *actionService) CreateAction(ctx *gin.Context) (string, error) {
	result := schemas.ActionResult{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return "", err
	}

	_, err = service.userService.GetUserInfos(tokenString)
	if err != nil {
		return "", err
	}

	serviceInfo := service.serviceService.FindByName(schemas.Github)
	newAction := schemas.Action{
		Name:        result.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: result.Description,
		ServiceId:   serviceInfo.Id,
		Options:     result.Options,
	}

	service.repository.Save(newAction)
	return "Action Created successfully", nil
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
