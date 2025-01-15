package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type WeatherService interface {
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
}

type weatherService struct {
	workflowRepository repository.WorkflowRepository
	mutex              sync.Mutex
}

func NewWeatherService(
	workflowRepository repository.WorkflowRepository,
) WeatherService {
	return &weatherService{
		workflowRepository: workflowRepository,
	}
}

func (service *weatherService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	switch name {
	case string(schemas.WeatherCurrentAction):
		return service.GetCurrentWeather
	default:
		return nil
	}
}

func (service *weatherService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	switch name {
	default:
		return nil
	}
}

func (service *weatherService) GetCurrentWeather(channel chan string, option string, workflowId uint64, actionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	apiKey := toolbox.GetInEnv("WEATHER_API_KEY")

	var actionData schemas.WeatherCurrentOptions
	err = json.Unmarshal([]byte(actionOption), &actionData)
	if err != nil {
		fmt.Println("Error parsing actionOption:", err)
		return
	}
	requestedUrl := "https://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + actionData.CityName + "&lang=" + actionData.LanguageCode
	request, err := http.NewRequest("GET", requestedUrl, nil)
	if err != nil {
		channel <- err.Error()
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

	var weatherResponse schemas.WeatherActionOptions
	bodyBytes, _ := io.ReadAll(response.Body)

	err = json.Unmarshal(bodyBytes, &weatherResponse)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		channel <- err.Error()
		return
	}
	switch actionData.CompareSign {
	case ">":
		if actionData.Temperature < weatherResponse.Current.Feelslike_c {
			service.UpdateWorkflowForAction(workflow, actionData)
			channel <- "Current weather"
			time.Sleep(30 * time.Second)
		}
	case "<":
		if actionData.Temperature > weatherResponse.Current.Feelslike_c {
			service.UpdateWorkflowForAction(workflow, actionData)
			channel <- "Current weather"
			time.Sleep(30 * time.Second)
		}
	case "=":
		{
			if actionData.Temperature == weatherResponse.Current.Feelslike_c {
				service.UpdateWorkflowForAction(workflow, actionData)
				channel <- "Current weather"
				time.Sleep(30 * time.Second)
			}
		}
	}
	channel <- "Current weather"
}

func (service *weatherService) UpdateWorkflowForAction(workflow schemas.Workflow, actionData schemas.WeatherCurrentOptions) {
	workflow.ReactionTrigger = true
	workflow.ActionOptions = toolbox.MustMarshal(actionData)
	service.workflowRepository.UpdateReactionTrigger(workflow)
}

func (service *weatherService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return nil
}
