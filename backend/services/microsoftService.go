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

type MicrosoftService interface {
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
	FindActionByName(name string) func(channel chan string, workflowId uint64, actionOption json.RawMessage)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage)
	AuthGetServiceAccessToken(code string, path string) (schemas.MicrosoftResponseToken, error)
}

type microsoftService struct {
	serviceToken       TokenService
	userService        UserService
	workflowRepository repository.WorkflowRepository
	serviceRepository  repository.ServiceRepository
	mutex              sync.Mutex
}

func NewMicrosoftService(
	serviceToken TokenService,
	userService UserService,
	workflowRepository repository.WorkflowRepository,
	serviceRepository repository.ServiceRepository,
) MicrosoftService {
	return &microsoftService{
		serviceToken:       serviceToken,
		userService:        userService,
		workflowRepository: workflowRepository,
		serviceRepository:  serviceRepository,
	}
}

func (service *microsoftService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return func(userInfos *schemas.ServicesUserInfos) {
		ctx := context.Background()

		url := "https://graph.microsoft.com/v1.0/me"
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return
		}

		req.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			// READ THE BODY
			bodyBytes, _ := io.ReadAll(resp.Body)
			fmt.Println("response body: ", string(bodyBytes))
			return
		}
		err = json.NewDecoder(resp.Body).Decode(&userInfos.MicrosoftUserInfos)
		if err != nil {
			return
		}
	}
}

func (service *microsoftService) AuthGetServiceAccessToken(code string, path string) (schemas.MicrosoftResponseToken, error) {
	clientId := toolbox.GetInEnv("MICROSOFT_CLIENT_ID")
	tenantId := toolbox.GetInEnv("MICROSOFT_TENANT_ID")
	appPort := toolbox.GetInEnv("FRONTEND_PORT")
	appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	redirectUri := appAdressHost + appPort + path
	apiUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantId)

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("code", code)
	data.Set("redirect_uri", redirectUri)
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return schemas.MicrosoftResponseToken{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{
		Timeout: time.Second * 45,
	}
	response, err := client.Do(req)
	fmt.Printf("response: %++v\n", response)
	if err != nil {
		return schemas.MicrosoftResponseToken{}, err
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Println("response body: ", string(bodyBytes))

	var result schemas.MicrosoftResponseToken
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return schemas.MicrosoftResponseToken{}, fmt.Errorf("unable to decode response because %w", err)
	}
	response.Body.Close()
	return result, nil
}

func (service *microsoftService) FindActionByName(name string) func(channel chan string, workflowId uint64, actionOption json.RawMessage) {
	switch name {
	case string(schemas.MicrosoftOutlookEventsAction):
		return service.GetOutlookEvents
	case string(schemas.MicrosoftTeamGroup):
		return service.ModifyTeamGroup
	default:
		return nil
	}
}

func (service *microsoftService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage) {
	switch name {
	case string(schemas.MicrosoftMailReaction):
		return service.SendMail
	default:
		return nil
	}
}

func (service *microsoftService) ModifyTeamGroup(channel chan string, workflowId uint64, actionOption json.RawMessage) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	workflow := service.workflowRepository.FindById(workflowId)
	user := service.userService.GetUserById(workflow.UserId)
	allTokens, err := service.serviceToken.GetTokenByUserId(user.Id)
	if err != nil {
		channel <- err.Error()
		return
	}

	options := schemas.MicrosoftTeamsChatResponse{}
	err = json.Unmarshal([]byte(actionOption), &options)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	url := "https://graph.microsoft.com/v1.0/me/chats/" + options.Id
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		channel <- err.Error()
		return
	}
	searchedService := service.serviceRepository.FindByName(schemas.Microsoft)
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

	result := schemas.MicrosoftTeamsChatResponse{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		channel <- err.Error()
		return
	}

	if options.IsOld {
		if options.LastUpdatedDateTime != result.LastUpdatedDateTime {
			workflow.ReactionTrigger = true
			service.workflowRepository.UpdateReactionTrigger(workflow)
			options.LastUpdatedDateTime = result.LastUpdatedDateTime
			workflow.ActionOptions = toolbox.RealObject(options)
			service.workflowRepository.Update(workflow)
		}
	} else {
		options.IsOld = true
		options.LastUpdatedDateTime = result.LastUpdatedDateTime
		workflow.ActionOptions = toolbox.RealObject(options)
		service.workflowRepository.Update(workflow)
	}
	channel <- "Action of modifying teams finished"
}

func (service *microsoftService) GetOutlookEvents(channel chan string, workflowId uint64, actionOption json.RawMessage) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	workflow := service.workflowRepository.FindById(workflowId)
	user := service.userService.GetUserById(workflow.UserId)
	allTokens, err := service.serviceToken.GetTokenByUserId(user.Id)
	if err != nil {
		channel <- err.Error()
		return
	}

	url := "https://graph.microsoft.com/v1.0/me/events"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		channel <- err.Error()
		return
	}
	searchedService := service.serviceRepository.FindByName(schemas.Microsoft)
	for _, token := range allTokens {
		if token.ServiceId == searchedService.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}

	options := schemas.MicrosoftOutlookEventsOptions{}
	err = json.Unmarshal([]byte(actionOption), &options)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	client := &http.Client{}
	request.Header.Set("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		channel <- err.Error()
		return
	}
	defer response.Body.Close()

	microsoftEventsSubjects := schemas.MicrosoftOutlookEventsResponse{}
	bodyBytes, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bodyBytes, &microsoftEventsSubjects)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		return
	}

	// fmt.Printf("VALUES : %s\n", string(bodyBytes))

	var chosenSubject *string
	for _, subject := range microsoftEventsSubjects.Value {
		if subject.Subject == options.Subject {
			chosenSubject = &subject.Subject
		}
	}

	if chosenSubject != nil {
		workflow.ReactionTrigger = true
		service.workflowRepository.Update(workflow)
	}
	channel <- "Finishing outlook action workflow"
}

func (service *microsoftService) SendMail(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage) {
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

	options := schemas.MicrosoftSendMailOptions{}
	err := json.Unmarshal([]byte(reactionOption), &options)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}
	url := "https://graph.microsoft.com/v1.0/me/sendMail"

	jsonData, err := json.Marshal(options)
	if err != nil {
		return
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	searchedService := service.serviceRepository.FindByName(schemas.Microsoft)
	for _, token := range accessToken {
		if token.ServiceId == searchedService.Id {
			request.Header.Set("Authorization", "Bearer "+token.Token)
		}
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	fmt.Printf("response: %++v\n", response)
	defer response.Body.Close()
	workflow.ReactionTrigger = false
	service.workflowRepository.UpdateReactionTrigger(workflow)
}
