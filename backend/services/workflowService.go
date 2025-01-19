package services

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"area51/repository"
	"area51/schemas"
	"area51/toolbox"
)

type WorkflowService interface {
	FindAll() []schemas.Workflow
	CreateWorkflow(ctx *gin.Context) (string, error)
	ActivateWorkflow(ctx *gin.Context) error
	InitWorkflow(workflowStartingPoint schemas.Workflow, githubServiceToken []schemas.ServiceToken, actionOption json.RawMessage, reactionOption json.RawMessage)
	ExistWorkflow(workflowId uint64) bool
	GetWorkflowByName(name string) schemas.Workflow
	GetWorkflowById(workflowId uint64) schemas.Workflow
	GetWorkflowsByUserId(userId uint64) []schemas.Workflow
	GetMostRecentReaction(ctx *gin.Context) ([]json.RawMessage, error)
	GetAllReactionsForAWorkflow(ctx *gin.Context) ([]json.RawMessage, error)
	DeleteWorkflow(ctx *gin.Context) error
	Delete(workflowId uint64) error
}

type workflowService struct {
	repository                  repository.WorkflowRepository
	userService                 UserService
	actionService               ActionService
	reactionService             ReactionService
	servicesService             ServicesService
	serviceToken                TokenService
	reactionResponseDataService ReactionResponseDataService
	googleRepository            repository.GoogleRepository
	githubRepository            repository.GithubRepository
}

func NewWorkflowService(
	repository repository.WorkflowRepository,
	userService UserService,
	actionService ActionService,
	reactionService ReactionService,
	servicesService ServicesService,
	serviceToken TokenService,
	reactionResponseDataService ReactionResponseDataService,
	googleRepository repository.GoogleRepository,
	githubRepository repository.GithubRepository,
) WorkflowService {
	return &workflowService{
		repository:                  repository,
		userService:                 userService,
		actionService:               actionService,
		reactionService:             reactionService,
		servicesService:             servicesService,
		serviceToken:                serviceToken,
		reactionResponseDataService: reactionResponseDataService,
		googleRepository:            googleRepository,
		githubRepository:            githubRepository,
	}
}

func (service *workflowService) FindAll() []schemas.Workflow {
	return service.repository.FindAll()
}

func (service *workflowService) CreateWorkflow(ctx *gin.Context) (string, error) {
	result := schemas.WorkflowResult{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return "", err
	}

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return "", err
	}

	workflowName := result.Name
	workflowValue := "1"
	if workflowName == "" {
		workflowName = "Workflow " + workflowValue
		for service.GetWorkflowByName(workflowName).Name != "" {
			workflowValueInt, _ := strconv.Atoi(workflowValue)
			workflowValueInt++
			workflowValue = strconv.Itoa(workflowValueInt)
			workflowName = "Workflow " + workflowValue
		}
		workflowValueInt, _ := strconv.Atoi(workflowValue)
		workflowValue = strconv.Itoa(workflowValueInt)
		workflowName = "Workflow " + workflowValue
	}

	serviceToken, _ := service.serviceToken.GetTokenByUserId(user.Id)
	newWorkflow := schemas.Workflow{
		UserId:          user.Id,
		User:            user,
		IsActive:        true,
		ActionId:        result.ActionId,
		ReactionId:      result.ReactionId,
		Action:          service.actionService.FindById(result.ActionId),
		ActionOptions:   result.ActionOption,
		Reaction:        service.reactionService.FindById(result.ReactionId),
		ReactionOptions: result.ReactionOption,
		Name:            workflowName,
	}
	actualWorkflow := service.repository.FindExistingWorkflow(newWorkflow)
	if actualWorkflow.Id != 0 {
		if actualWorkflow.IsActive {
			return "Workflow already exists and is active", nil
		} else {
			return "Workflow already exists and is not active", nil
		}
	}
	fmt.Println(newWorkflow)
	workflowId, err := service.repository.SaveWorkflow(newWorkflow)
	if err != nil {
		return "", err
	}

	newWorkflow.Id = workflowId
	service.InitWorkflow(newWorkflow, serviceToken, result.ActionOption, result.ReactionOption)
	return "Workflow Created succesfully", nil

}

func (service *workflowService) ActivateWorkflow(ctx *gin.Context) error {
	var result schemas.WorkflowActivate
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return err
	}
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) <= len("Bearer ") {
		return fmt.Errorf("no authorization header found")
	}
	tokenString := authHeader[len("Bearer "):]

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return err
	}
	workflow, err := service.repository.FindByIds(result.WorkflowId)
	if err != nil {
		return err
	}
	newWorkflow := schemas.Workflow{
		Id:              workflow.Id,
		UserId:          user.Id,
		IsActive:        result.WorflowState,
		ReactionTrigger: result.WorflowState,
		ActionOptions:   workflow.ActionOptions,
		ReactionOptions: workflow.ReactionOptions,
	}
	service.repository.UpdateActiveStatus(newWorkflow)
	service.repository.UpdateReactionTrigger(newWorkflow)
	return nil
}

func (service *workflowService) InitWorkflow(workflowStartingPoint schemas.Workflow, githubServiceToken []schemas.ServiceToken, actionOption json.RawMessage, reactionOption json.RawMessage) {
	workflowChannel := make(chan string)
	go service.WorkflowActionChannel(workflowStartingPoint, workflowChannel, actionOption)
	go service.WorkflowReactionChannel(workflowStartingPoint, workflowChannel, githubServiceToken, reactionOption)
}

func (service *workflowService) WorkflowActionChannel(workflowStartingPoint schemas.Workflow, channel chan string, actionOption json.RawMessage) {
	go func(workflowStartingPoint schemas.Workflow, channel chan string) {
		fmt.Println("Start of WorkflowActionChannel")
		for service.ExistWorkflow(workflowStartingPoint.Id) {
			workflow, err := service.repository.FindByIds(workflowStartingPoint.Id)
			if err != nil {
				fmt.Println("Error")
				return
			}
			action := service.servicesService.FindActionByName(workflow.Action.Name)
			if action == nil {
				fmt.Println("Action not found", workflow.Action.Name)
				return
			}
			if workflow.IsActive {
				action(channel, workflow.Id, actionOption)
			}
			time.Sleep(30 * time.Second)
		}
		fmt.Println("Clear")
		channel <- "Workflow finished"
	}(workflowStartingPoint, channel)
}

func (service *workflowService) WorkflowReactionChannel(workflowStartingPoint schemas.Workflow, channel chan string, githubServiceToken []schemas.ServiceToken, reactionOption json.RawMessage) {
	go func(workflowStartingPoint schemas.Workflow, channel chan string) {
		for service.ExistWorkflow(workflowStartingPoint.Id) {
			workflow, err := service.repository.FindByIds(workflowStartingPoint.Id)
			if err != nil {
				fmt.Println("Error")
				return
			}

			reaction := service.servicesService.FindReactionByName(workflow.Reaction.Name)
			if reaction == nil {
				fmt.Println("Reaction not found")
				return
			}

			if workflow.IsActive {
				result := <-channel
				reaction(channel, workflow.Id, githubServiceToken, reactionOption)
				fmt.Printf("result value: %+v\n", result)
			}
			time.Sleep(30 * time.Second)
		}
	}(workflowStartingPoint, channel)
}

func (service *workflowService) ExistWorkflow(workflowId uint64) bool {
	_, err := service.repository.FindByIds(workflowId)
	return err == nil
}

func (service *workflowService) GetWorkflowByName(name string) schemas.Workflow {
	return service.repository.FindByWorkflowName(name)
}

func (service *workflowService) GetWorkflowById(workflowId uint64) schemas.Workflow {
	return service.repository.FindById(workflowId)
}

func (service *workflowService) GetWorkflowsByUserId(userId uint64) []schemas.Workflow {
	return service.repository.FindByUserId(userId)
}

func (service *workflowService) Delete(workflowId uint64) error {
	return service.repository.Delete(workflowId)
}

func (service *workflowService) GetMostRecentReaction(ctx *gin.Context) ([]json.RawMessage, error) {
	var result schemas.ReactionOutputJson
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	_, err = toolbox.GetBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	workflow := service.repository.FindById(result.WorkflowId)
	if workflow.Id == 0 {
		return nil, fmt.Errorf("workflow not found")
	}
	reactionResponse := []json.RawMessage{}
	reactionResponseData := service.reactionResponseDataService.FindByWorkflowId(workflow.Id)
	tmp := schemas.ReactionResponseData{}
	for _, data := range reactionResponseData {
		if !data.CreatedAt.Before(tmp.CreatedAt) {
			tmp = data
		}
	}
	reactionResponse = append(reactionResponse, tmp.ApiResponse)
	return reactionResponse, nil
}

func (service *workflowService) DeleteWorkflow(ctx *gin.Context) error {
	var result schemas.WorkflowJson
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return err
	}
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) <= len("Bearer ") {
		return fmt.Errorf("no authorization header found")
	}
	tokenString := authHeader[len("Bearer "):]

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return err
	}
	workflow := service.repository.FindAll()
	for _, wf := range workflow {
		if wf.Name == result.Name && wf.UserId == user.Id && wf.ReactionId == result.ReactionId && wf.ActionId == result.ActionId {
			actualReactionData := service.reactionResponseDataService.FindByWorkflowId(wf.Id)
			for _, data := range actualReactionData {
				service.reactionResponseDataService.Delete(data)
			}
			err := service.repository.Delete(wf.Id)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (service *workflowService) GetAllReactionsForAWorkflow(ctx *gin.Context) ([]json.RawMessage, error) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) <= len("Bearer ") {
		return nil, fmt.Errorf("no authorization header found")
	}
	tokenString := authHeader[len("Bearer "):]

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return nil, err
	}
	workflow := service.repository.FindByUserId(user.Id)
	var reactionResponse []json.RawMessage
	for _, wf := range workflow {
		if wf.UserId == user.Id {
			reactionResponseData := service.reactionResponseDataService.FindByWorkflowId(wf.Id)
			for _, data := range reactionResponseData {
				reactionResponse = append(reactionResponse, data.ApiResponse)
			}
		}
	}
	return reactionResponse, nil
}
