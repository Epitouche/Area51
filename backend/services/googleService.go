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
	mutex              sync.Mutex
}

func NewGoogleService(
	serviceToken TokenService,
	userService UserService,
	workflowRepository repository.WorkflowRepository,
	serviceRepository repository.ServiceRepository,
) GoogleService {
	return &googleService{
		serviceToken:       serviceToken,
		userService:        userService,
		workflowRepository: workflowRepository,
		serviceRepository:  serviceRepository,
	}
}

func (service *googleService) AuthGetServiceAccessToken(code string, path string) (schemas.GoogleResponseToken, error) {
	clientId := toolbox.GetInEnv("GOOGLE_CLIENT_ID")
	clientSecret := toolbox.GetInEnv("GOOGLE_SECRET")
	// appPort := toolbox.GetInEnv("FRONTEND_PORT")
	// appAdressHost := toolbox.GetInEnv("APP_HOST_ADDRESS")

	//! TEMPORARY
	// redirectUri := appAdressHost + appPort + path
	redirectUri := "http://localhost:8081/callback"
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
	return nil
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
		fmt.Println(err)
		time.Sleep(30 * time.Second)
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
			fmt.Printf("TOKEN %s\n", token.Token)
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
	bodyBytes, _ := io.ReadAll(response.Body)
	// fmt.Printf("Value: %s\n", string(bodyBytes))
}

// resultSizeEstimate
