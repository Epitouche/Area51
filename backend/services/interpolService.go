package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type InterpolService interface {
	FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
}

type interpolService struct {
	workflowRepository repository.WorkflowRepository
	mutex          sync.Mutex
}

func NewInterpolService(
	workflowRepository repository.WorkflowRepository,
) InterpolService {
	return &interpolService{
		workflowRepository: workflowRepository,
	}
}

func (service *interpolService) FindActionByName(name string) func(channel chan string, option string, workflowId uint64, actionOption string) {
	switch name {
	case string(schemas.InterpolNewNotices):
		return service.NewNotices
	default:
		return nil
	}
}

func (service *interpolService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption string) {
	switch name {
	default:
		return nil
	}
}

func (service *interpolService) NewNotices(channel chan string, option string, workflowId uint64, actionOption string) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	request, err := http.NewRequest("GET", "https://ws-public.interpol.int/notices/v1/red", nil)
	if err != nil {
		fmt.Printf("unable to create request because: %s", err)
		time.Sleep(30 * time.Second)
		return
	}
    request.Header.Set("Accept", "application/json")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cache-Control", "no-cache")
    client := &http.Client{}
	
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	defer response.Body.Close()
	result := schemas.InterpolRedNoticesInfos{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Reponse err")
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}

	options := schemas.InterpolActionOption{}
	if actionOption != "" {
		err = json.NewDecoder(strings.NewReader(actionOption)).Decode(&options)
		if err != nil {
			fmt.Println("Options err")
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			return
		}
	}

	workflow, err := service.workflowRepository.FindByIds(workflowId)
	if err != nil {
		fmt.Println(err)
		time.Sleep(30 * time.Second)
		return
	}
	if options.IsOld {
		if result.Total != options.Total {
			options.Total = result.Total
			workflow.ReactionTrigger = true
			workflow.ActionOptions = toolbox.MustMarshal(options)
			service.workflowRepository.Update(workflow)
		}
	} else {
		options.Total = result.Total
		options.IsOld = true
		workflow.ActionOptions = toolbox.MustMarshal(options)
		service.workflowRepository.Update(workflow)
	}
	channel <- "Action worlflow done"
	time.Sleep(30 * time.Second)
}

func (service *interpolService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return nil
}
