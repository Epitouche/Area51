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
				Name:        string(schemas.GithubReactionListComments),
				Description: "List all comments of a repository",
				ServiceId:   serviceService.FindByName(schemas.Github).Id,
				Options: toolbox.RealObject(schemas.GithubListAllReviewCommentsOptions{
					Owner: "my github username",
					Repo:  "name of the repository",
				}),
			},
			{
				Name:        string(schemas.SpotifyAddTrackReaction),
				Description: "Add a track to a playlist",
				ServiceId:   serviceService.FindByName(schemas.Spotify).Id,
				Options: toolbox.RealObject(schemas.SpotifyReactionOptions{
					PlaylistURL: "https://open.spotify.com/playlist/37i9dQZF1DXcBWIGoYBM5M",
					TrackURL:    "https://open.spotify.com/track/4PTG3Z6ehGkBFwjybzWkR8",
				}),
			},
			{
				Name:        string(schemas.GoogleCreateEventReaction),
				Description: "Create an event in Google Calendar",
				ServiceId:   serviceService.FindByName(schemas.Google).Id,
				Options: toolbox.RealObject(schemas.GoogleCalendarOptionsSchema{
					CalendarId: "your address email",
					CalendarCorpus: schemas.GoogleCalendarCorpusOptionsSchema{
						Summary:     "Réunion",
						Description: "on va parler de l'avenir",
						Location:    "Osaka",
						Start: schemas.GoogleCalendarCorpusOptionsTimeStartSchema{
							StartDateTime: "2025-01-15T10:00:00.0000000",
							StartTimeZone: "Europe/Paris",
						},
						End: schemas.GoogleCalendarCorpusOptionsTimeEndSchema{
							EndDateTime: "2025-01-15T10:00:00.0000000",
							EndTimeZone: "Europe/Paris",
						},
						Attendees: schemas.GoogleCalendarCorpusOptionsAttendees{
							Email: "my.email@gmail.com",
						},
					},
				}),
			},
			{
				Name:        "get_red_notices",
				Description: "Detect a change on a specific notice",
				ServiceId:   serviceService.FindByName(schemas.Interpol).Id,
				Options: toolbox.RealObject(schemas.InterpolReactionOptionInfos{
					FirstName: "Sylvain",
					LastName:  "téun",
				}),
			},
			{
				Name:        "get_yellow_notices",
				Description: "Detect a change on a specific notice",
				ServiceId:   serviceService.FindByName(schemas.Interpol).Id,
				Options: toolbox.RealObject(schemas.InterpolReactionOptionInfos{
					FirstName: "Michel",
					LastName:  "Levoisin",
				}),
			},
			{
				Name:        "get_un_notices",
				Description: "Detect a change on a specific notice",
				ServiceId:   serviceService.FindByName(schemas.Interpol).Id,
				Options: toolbox.RealObject(schemas.InterpolReactionOptionInfos{
					FirstName: "Gérard",
					LastName:  "Auplacard",
				}),
			},
			{
				Name:        string(schemas.WeatherCurrentReaction),
				Description: "Get the current weather of a city",
				ServiceId:   serviceService.FindByName(schemas.Weather).Id,
				Options: toolbox.RealObject(schemas.WeatherCurrentReactionOptions{
					CityName:     "Bordeaux",
					LanguageCode: "FR",
				}),
			},
			{
				Name:        string(schemas.MicrosoftMailReaction),
				Description: "Send an email",
				ServiceId:   serviceService.FindByName(schemas.Microsoft).Id,
				Options: toolbox.RealObject(schemas.MicrosoftSendMailOptionsSchema{
					Message: schemas.MicrosoftSendMailMainMessageOptionsSchema{
						Subject: "We are going to Chicoutimi ?",
						Body: schemas.MicrosoftSendMailBodyOptions{
							ContentType: "Text",
							Content:     "This email is to confirm our trip to Chicoutimi",
						},
						Address: "my.email@gmail.com",
					},
					SaveToSentItems: "true / false",
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
			Options:     oneReaction.Options,
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
