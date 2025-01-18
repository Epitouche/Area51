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
	workflowRepository          repository.WorkflowRepository
	userService                 UserService
	reactionResponseDataService ReactionResponseDataService
	mutex                       sync.Mutex
}

func NewWeatherService(
	workflowRepository repository.WorkflowRepository,
	userService UserService,
	reactionResponseDataService ReactionResponseDataService,
) WeatherService {
	return &weatherService{
		workflowRepository:          workflowRepository,
		userService:                 userService,
		reactionResponseDataService: reactionResponseDataService,
	}
}

func (service *weatherService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	switch name {
	case string(schemas.WeatherCurrentAction):
		return service.VerifyFeelingTemperature
	case string(schemas.WeatherTimeAction):
		return service.SunriseEvents
	default:
		return nil
	}
}

func (service *weatherService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	switch name {
	case string(schemas.WeatherCurrentReaction):
		return service.GetCurrentWeather
	default:
		return nil
	}
}

func (service *weatherService) VerifyFeelingTemperature(channel chan string, option string, workflowId uint64, actionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
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
			// time.Sleep(30 * time.Second)
		}
	case "<":
		if actionData.Temperature > weatherResponse.Current.Feelslike_c {
			service.UpdateWorkflowForAction(workflow, actionData)
			channel <- "Current weather"
			// time.Sleep(30 * time.Second)
		}
	case "=":
		{
			if actionData.Temperature == weatherResponse.Current.Feelslike_c {
				service.UpdateWorkflowForAction(workflow, actionData)
				channel <- "Current weather"
				// time.Sleep(30 * time.Second)
			}
		}
	}
	// channel <- "Current weather"
	// time.Sleep(30 * time.Second)
}

func (service *weatherService) GetCurrentWeather(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

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

	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
		return
	}

	apiKey := toolbox.GetInEnv("WEATHER_API_KEY")

	var actionData schemas.WeatherCurrentReactionOptions
	err = json.Unmarshal([]byte(reactionOption), &actionData)
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
	var weatherResponse schemas.WeatherReactionOptions
	bodyBytes, _ := io.ReadAll(response.Body)

	err = json.Unmarshal(bodyBytes, &weatherResponse)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		channel <- err.Error()
		return
	}
	savedResult := schemas.ReactionResponseData{
		WorkflowId:  workflowId,
		Workflow:    workflow,
		ApiResponse: json.RawMessage{},
	}
	jsonValue, err := json.Marshal(weatherResponse)
	if err != nil {
		fmt.Println("Error marshalling response:", err)
		return
	}
	savedResult.ApiResponse = jsonValue
	service.reactionResponseDataService.Save(savedResult)
	workflow.ReactionTrigger = false
	service.workflowRepository.UpdateReactionTrigger(workflow)
}

func (service *weatherService) UpdateWorkflowForAction(workflow schemas.Workflow, actionData schemas.WeatherCurrentOptions) {
	workflow.ReactionTrigger = true
	workflow.ActionOptions = toolbox.MustMarshal(actionData)
	service.workflowRepository.Update(workflow)
}

func (service *weatherService) SunriseEvents(channel chan string, option string, workflowId uint64, actionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	apiKey := toolbox.GetInEnv("WEATHER_API_KEY")

	var actionData schemas.WeatherSpecificTimeOption
	err = json.Unmarshal([]byte(actionOption), &actionData)
	if err != nil {
		fmt.Println("Error parsing actionOption:", err)
		return
	}
	requestedUrl := "https://api.weatherapi.com/v1/astronomy.json?key=" + apiKey + "&q=" + actionData.CityName + "&dt=" + actionData.DateTime
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

	var weatherResponse schemas.WeatherSpecificTimeInfos
	bodyBytes, _ := io.ReadAll(response.Body)
	fmt.Printf("Body: %s\n", string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &weatherResponse)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		channel <- err.Error()
		return
	}
	actualTime := time.Now().Format("15:04")

	realTimeValue := strToTime(actualTime)
	realTimeOption := strToTime(weatherResponse.Astronomy.Astro.Sunrise[0:5])
	if realTimeOption == realTimeValue {
		workflow.ReactionTrigger = true
		workflow.ActionOptions = toolbox.MustMarshal(actionData)
		service.workflowRepository.Update(workflow)
	} else {
		workflow.ReactionTrigger = false
		workflow.ActionOptions = toolbox.MustMarshal(actionData)
		service.workflowRepository.UpdateReactionTrigger(workflow)
		fmt.Println("Time is not equal")
		return
	}

	channel <- "Specific Time action done"
}

func strToTime(s string) time.Time {
	t, _ := time.Parse("15:04", s)
	return t
}

func (service *weatherService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return nil
}
