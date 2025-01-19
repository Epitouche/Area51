package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type GoogleService interface {
	AuthGetServiceAccessToken(code string, path string) (schemas.GoogleResponseToken, error)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
	FindActionByName(name string) func(channel chan string, workflowId uint64, actionOption json.RawMessage)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage)
}

type googleService struct {
	serviceToken       TokenService
	userService        UserService
	workflowRepository repository.WorkflowRepository
	serviceRepository  repository.ServiceRepository
	googleRepository   repository.GoogleRepository
	mutex              sync.Mutex
}

func NewGoogleService(
	serviceToken TokenService,
	userService UserService,
	workflowRepository repository.WorkflowRepository,
	serviceRepository repository.ServiceRepository,
	googleRepository repository.GoogleRepository,
) GoogleService {
	return &googleService{
		serviceToken:       serviceToken,
		userService:        userService,
		workflowRepository: workflowRepository,
		serviceRepository:  serviceRepository,
		googleRepository:   googleRepository,
	}
}

func (service *googleService) AuthGetServiceAccessToken(code string, path string) (schemas.GoogleResponseToken, error) {
	clientId := toolbox.GetInEnv("GOOGLE_CLIENT_ID")
	clientSecret := toolbox.GetInEnv("GOOGLE_SECRET")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	redirectUri := appAdressHost + appPort + path
	apiUrl := "https://oauth2.googleapis.com/token"
	decodedCode, err := url.QueryUnescape(code)
	if err != nil {
		return schemas.GoogleResponseToken{}, err
	}

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("code", decodedCode)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return schemas.GoogleResponseToken{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		Timeout: time.Second * 45,
	}
	response, err := client.Do(req)
	if err != nil {
		return schemas.GoogleResponseToken{}, err
	}
	bodyBytes, _ := io.ReadAll(response.Body)

	var result schemas.GoogleResponseToken
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return schemas.GoogleResponseToken{}, fmt.Errorf("unable to decode response because %w", err)
	}
	response.Body.Close()
	return result, nil
}

func (service *googleService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return func(userInfos *schemas.ServicesUserInfos) {
		ctx := context.Background()

		url := "https://www.googleapis.com/oauth2/v1/userinfo"
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return
		}

		request.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: received status code %d\n", response.StatusCode)
			return
		}

		err = json.NewDecoder(response.Body).Decode(&userInfos.GoogleUserInfos)
		if err != nil {
			return
		}
	}
}

func (service *googleService) FindActionByName(name string) func(channel chan string, workflowId uint64, actionOption json.RawMessage) {
	switch name {
	case string(schemas.GoogleGetEmailAction):
		return service.GetEmailAction
	default:
		return nil
	}
}

func (service *googleService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage) {
	switch name {
	case string(schemas.GoogleCreateEventReaction):
		return service.CreateEventReaction
	default:
		return nil
	}
}

func (service *googleService) GetEmailAction(channel chan string, workflowId uint64, actionOption json.RawMessage) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	workflow := service.workflowRepository.FindById(workflowId)
	user := service.userService.GetUserById(workflow.UserId)
	allTokens, err := service.serviceToken.GetTokenByUserId(user.Id)
	if err != nil {
		channel <- err.Error()
		return
	}

	options := schemas.GoogleActionOptions{}
	err = json.Unmarshal([](byte)(actionOption), &options)
	if err != nil {
		fmt.Println("Error parsing actionOption:", err)
		return
	}

	url := "https://www.googleapis.com/gmail/v1/users/me/messages?labelIds=" + options.Label
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		channel <- err.Error()
		return
	}
	searchedService := service.serviceRepository.FindByName(schemas.Google)
	for _, token := range allTokens {
		if token.ServiceId == searchedService.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}
	client := &http.Client{}
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		channel <- err.Error()
		return
	}
	defer response.Body.Close()
	time.Sleep(10 * time.Millisecond)
	googleOption := schemas.GoogleActionOptionsInfo{}
	bodyBytes, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bodyBytes, &googleOption)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		// 	channel <- err.Error()
		return
	}
	existingRecords := map[string]interface{}{}

	if string(workflow.Utils) != "" {
		err = json.Unmarshal([]byte(workflow.Utils), &existingRecords)
		if err != nil {
			fmt.Println("Error unmarshalling existingRecords:", err)
			return
		}
	}

	if existingRecords["ResultSizeEstimate"] == nil {
		existingRecords["ResultSizeEstimate"] = 0
		jsonData, err := json.Marshal(existingRecords)
		if err != nil {
			fmt.Println("Error marshalling existingRecords:", err)
			return
		}
		workflow.Utils = jsonData
		service.workflowRepository.Update(workflow)
	}
	var ResultSizeEstimate int
	switch v := existingRecords["ResultSizeEstimate"].(type) {
	case float64:
		ResultSizeEstimate = int(v)
	case int:
		ResultSizeEstimate = v
	default:
		fmt.Println("Error asserting NumPR to int or float64")
		return
	}

	existingRecords["ResultSizeEstimate"] = googleOption.ResultSizeEstimate
	jsonData, err := json.Marshal(existingRecords)
	if err != nil {
		fmt.Println("Error marshalling existingRecords:", err)
		return
	}
	workflow.Utils = jsonData
	if ResultSizeEstimate < googleOption.ResultSizeEstimate {
		workflow.ReactionTrigger = true
		service.workflowRepository.UpdateReactionTrigger(workflow)
	}
	service.workflowRepository.UpdateUtils(workflow)
	channel <- "Emails fetched"
}

func (service *googleService) CreateEventReaction(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	// Iterate over all tokens
	for _, token := range accessToken {
		actualUser := service.userService.GetUserById(token.UserId)
		if token.UserId == actualUser.Id {
			actualWorkflow := service.workflowRepository.FindByUserId(actualUser.Id)
			for _, workflow := range actualWorkflow {
				if workflow.Id == workflowId {
					if !workflow.ReactionTrigger {
						fmt.Println("Trigger is already false, skipping reaction.")
						return
					}
				}
			}
		}
	}

	workflow := service.workflowRepository.FindById(workflowId)

	url := "https://www.googleapis.com/calendar/v3/users/me/calendarList"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	searchedService := service.serviceRepository.FindByName(schemas.Google)
	for _, token := range accessToken {
		if token.ServiceId == searchedService.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}

	options := schemas.GoogleCalendarOptionsSchema{}
	err = json.Unmarshal([]byte(reactionOption), &options)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}
	trueOptions := schemas.GoogleCalendarOptions{
		CalendarId: options.CalendarId,
		CalendarCorpus: schemas.GoogleCalendarCorpusOptions{
			Summary:     options.CalendarCorpus.Summary,
			Description: options.CalendarCorpus.Description,
			Location:    options.CalendarCorpus.Location,
			Start: schemas.GoogleCalendarCorpusOptionsTime{
				DateTime: options.CalendarCorpus.Start.StartDateTime,
				TimeZone: options.CalendarCorpus.Start.StartTimeZone,
			},
			End: schemas.GoogleCalendarCorpusOptionsTime{
				DateTime: options.CalendarCorpus.End.EndDateTime,
				TimeZone: options.CalendarCorpus.End.EndTimeZone,
			},
			Attendees: options.CalendarCorpus.Attendees,
		},
	}

	client := &http.Client{}
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	time.Sleep(10 * time.Millisecond)
	googleCalendarIds := schemas.GoogleCalendarResponse{}
	bodyBytes, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bodyBytes, &googleCalendarIds)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		// 	channel <- err.Error()
		return
	}
	var wantedCaledarId string
	for _, calendar := range googleCalendarIds.Items {
		if calendar.Id == trueOptions.CalendarId {
			wantedCaledarId = calendar.Id
		}
	}
	response.Body.Close()

	urlToCreateEvent := "https://www.googleapis.com/calendar/v3/calendars/" + wantedCaledarId + "/events"

	jsonData, err := json.Marshal(trueOptions.CalendarCorpus)
	if err != nil {
		return
	}

	request, err = http.NewRequest("POST", urlToCreateEvent, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	for _, token := range accessToken {
		if token.ServiceId == searchedService.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}
	request.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	workflow.ReactionTrigger = false
	service.workflowRepository.UpdateReactionTrigger(workflow)

}
