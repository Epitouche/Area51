package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"area51/repository"
	"area51/schemas"
)

type InterpolService interface {
	FindActionByName(name string) func(channel chan string, workflowId uint64, actionOption json.RawMessage)
	FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage)
	GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos)
}

type interpolService struct {
	workflowRepository          repository.WorkflowRepository
	reactionRepository          repository.ReactionRepository
	userService                 UserService
	reactionResponseDataService ReactionResponseDataService
	mutex                       sync.Mutex
}

func NewInterpolService(
	workflowRepository repository.WorkflowRepository,
	reactionRepository repository.ReactionRepository,
	userService UserService,
	reactionResponseDataService ReactionResponseDataService,
) InterpolService {
	return &interpolService{
		workflowRepository:          workflowRepository,
		reactionRepository:          reactionRepository,
		userService:                 userService,
		reactionResponseDataService: reactionResponseDataService,
	}
}

func (service *interpolService) FindActionByName(name string) func(channel chan string, workflowId uint64, actionOption json.RawMessage) {
	switch name {
	default:
		return nil
	}
}

func (service *interpolService) FindReactionByName(name string) func(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage) {
	switch name {
	case string(schemas.InterpolGetRedNotices):
		return service.GetNotices
	case string(schemas.InterpolGetYellowNotices):
		return service.GetNotices
	case string(schemas.InterpolGetUNNotices):
		return service.GetNotices
	default:
		return nil
	}
}

func (service *interpolService) GetNotices(channel chan string, workflowId uint64, accessToken []schemas.ServiceToken, reactionOption json.RawMessage) {
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

	workflow := service.workflowRepository.FindById(workflowId)
	reaction := service.reactionRepository.FindById(workflow.ReactionId)
	noticeType := ""
	switch reaction.Name {
	case string(schemas.InterpolGetRedNotices):
		noticeType = "red"
	case string(schemas.InterpolGetYellowNotices):
		noticeType = "yellow"
	case string(schemas.InterpolGetUNNotices):
		noticeType = "un"
	}
	options := schemas.InterpolReactionOption{}
	err := json.Unmarshal([]byte(reaction.Options), &options)
	if err != nil {
		fmt.Println("Error ->", err)
		return
	}

	request, err := http.NewRequest("GET", "https://ws-public.interpol.int/notices/v1/"+noticeType+"?forename="+options.FirstName+"&name="+options.LastName, nil)
	if err != nil {
		fmt.Printf("unable to create request because: %s", err)
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
		return
	}

	result := schemas.InterpolNoticesList{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return
	}
	savedResult := schemas.ReactionResponseData{
		WorkflowId:  workflowId,
		ApiResponse: json.RawMessage{},
	}
	jsonValue, err := json.Marshal(result.Embedded.Notices)
	if err != nil {
		fmt.Println("Error marshalling response:", err)
		return
	}
	savedResult.ApiResponse = jsonValue
	service.reactionResponseDataService.Save(savedResult)
	workflow.ReactionTrigger = false
	service.workflowRepository.UpdateReactionTrigger(workflow)
}

func (service *interpolService) GetUserInfosByToken(accessToken string, serviceName schemas.ServiceName) func(*schemas.ServicesUserInfos) {
	return nil
}
