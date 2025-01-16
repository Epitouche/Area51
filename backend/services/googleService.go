package services

import (
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
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
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

		// fmt.Printf("response: %++v\n", response)
		// Read the response body
		// bodyBytes, _ := io.ReadAll(response.Body)
		// fmt.Println("response body: ", string(bodyBytes))

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

func (service *googleService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	switch name {
	case string(schemas.GoogleGetEmailAction):
		return service.GetEmailAction
	default:
		return nil
	}
}

func (service *googleService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	switch name {
	case string(schemas.GoogleCreateEventReaction):
		return service.CreateEventReaction
	default:
		return nil
	}
}

func (service *googleService) GetEmailAction(channel chan string, option string, workflowId uint64, actionOption string) {
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
	err = json.NewDecoder(strings.NewReader(workflow.ActionOptions)).Decode(&options)
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
	existingRecords := service.googleRepository.FindByWorkflowId(workflowId)
	if existingRecords.UserId == 0 {
		service.googleRepository.Save(schemas.GoogleActionResponse{
			User:               user,
			UserId:             user.Id,
			Worflow:            workflow,
			WorkflowId:         workflowId,
			ResultSizeEstimate: 0,
		})
	}
	if existingRecords.ResultSizeEstimate < googleOption.ResultSizeEstimate {
		workflow.ReactionTrigger = true
		service.workflowRepository.Update(workflow)
	}
	actualRecords := service.googleRepository.FindByWorkflowId(workflowId)
	actualRecords.ResultSizeEstimate = googleOption.ResultSizeEstimate
	service.googleRepository.UpdateNumEmails(actualRecords)
	channel <- "Emails fetched"
}

func (service *googleService) CreateEventReaction(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	// Get the current calendars of the current user
	url := "https://www.googleapis.com/calendar/v3/users/me/calendarList"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	fmt.Printf("Request created: %++v\n", request)

	searchedService := service.serviceRepository.FindByName(schemas.Google)
	for _, token := range accessToken {
		if token.ServiceId == searchedService.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}

	client := &http.Client{}
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	defer response.Body.Close()
	time.Sleep(10 * time.Millisecond)
	googleCalendarIds := schemas.GoogleCalendarResponse{}
	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Printf("Value of bodyBytes: %s\n", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &googleCalendarIds)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		// 	channel <- err.Error()
		return
	}
	fmt.Printf("Value of googleCalendarIds: %++v\n", googleCalendarIds)
	panic("not implemented\n")
	// err = json.Unmarshal(bodyBytes, &googleOption)
	// if err != nil {
	// 	fmt.Printf("Error: %s\n", err)
	// 	channel <- err.Error()
	// 	return
	// }

}

// Value of bodyBytes: {
// 	"kind": "calendar#calendarList",
// 	"etag": "\"p32s8v0fbmejog0o\"",
// 	"nextSyncToken": "CLiPgeuzp4gDEhVraWtpMjkzMy5rdEBnbWFpbC5jb20=",
// 	"items": [
// 	 {
// 	  "kind": "calendar#calendarListEntry",
// 	  "etag": "\"1667657290185000\"",
// 	  "id": "kiki2933.kt@gmail.com",
// 	  "summary": "kiki2933.kt@gmail.com",
// 	  "timeZone": "UTC",
// 	  "colorId": "14",
// 	  "backgroundColor": "#9fe1e7",
// 	  "foregroundColor": "#000000",
// 	  "selected": true,
// 	  "accessRole": "owner",
// 	  "defaultReminders": [
// 	   {
// 		"method": "popup",
// 		"minutes": 30
// 	   }
// 	  ],
// 	  "notificationSettings": {
// 	   "notifications": [
// 		{
// 		 "type": "eventCreation",
// 		 "method": "email"
// 		},
// 		{
// 		 "type": "eventChange",
// 		 "method": "email"
// 		},
// 		{
// 		 "type": "eventCancellation",
// 		 "method": "email"
// 		},
// 		{
// 		 "type": "eventResponse",
// 		 "method": "email"
// 		}
// 	   ]
// 	  },
// 	  "primary": true,
// 	  "conferenceProperties": {
// 	   "allowedConferenceSolutionTypes": [
// 		"hangoutsMeet"
// 	   ]
// 	  }
// 	 },
// 	 {
// 	  "kind": "calendar#calendarListEntry",
// 	  "etag": "\"1667657291887000\"",
// 	  "id": "td3gu9qvbokpfpas267a4qt2j4@group.calendar.google.com",
// 	  "summary": "TROUVE KILLIAN",
// 	  "timeZone": "UTC",
// 	  "colorId": "3",
// 	  "backgroundColor": "#f83a22",
// 	  "foregroundColor": "#000000",
// 	  "selected": true,
// 	  "accessRole": "owner",
// 	  "defaultReminders": [],
// 	  "conferenceProperties": {
// 	   "allowedConferenceSolutionTypes": [
// 		"hangoutsMeet"
// 	   ]
// 	  }
// 	 },
// 	 {
// 	  "kind": "calendar#calendarListEntry",
// 	  "etag": "\"1705357991846000\"",
// 	  "id": "fr.french#holiday@group.v.calendar.google.com",
// 	  "summary": "Jours fériés et autres fêtes en France",
// 	  "description": "Jours fériés et fêtes légales en France",
// 	  "timeZone": "UTC",
// 	  "colorId": "8",
// 	  "backgroundColor": "#16a765",
// 	  "foregroundColor": "#000000",
// 	  "selected": true,
// 	  "accessRole": "reader",
// 	  "defaultReminders": [],
// 	  "conferenceProperties": {
// 	   "allowedConferenceSolutionTypes": [
// 		"hangoutsMeet"
// 	   ]
// 	  }
// 	 },
// 	 {
// 	  "kind": "calendar#calendarListEntry",
// 	  "etag": "\"1705357992518000\"",
// 	  "id": "addressbook#contacts@group.v.calendar.google.com",
// 	  "summary": "Anniversaires",
// 	  "description": "Affiche dans Google Contacts les dates d'anniversaire et autres dates importantes de vos contacts.",
// 	  "timeZone": "UTC",
// 	  "colorId": "13",
// 	  "backgroundColor": "#92e1c0",
// 	  "foregroundColor": "#000000",
// 	  "selected": true,
// 	  "accessRole": "reader",
// 	  "defaultReminders": [],
// 	  "conferenceProperties": {
// 	   "allowedConferenceSolutionTypes": [
// 		"hangoutsMeet"
// 	   ]
// 	  }
// 	 }
// 	]
//    }

//    panic: not implemented
