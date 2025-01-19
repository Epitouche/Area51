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
	Update(ctx *gin.Context) error
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
	err := ctx.ShouldBind(&result)
	if err != nil {
		return "", schemas.ErrorBadParameter
	}

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return "", err
	}

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return "", schemas.ErrUserNotFound
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

	reaction := service.reactionService.FindById(result.ReactionId)
	action := service.actionService.FindById(result.ActionId)
	if reaction.Id == 0 {
		return "", schemas.ErrReactionNotFound
	}
	if action.Id == 0 {
		return "", schemas.ErrActionNotFound
	}
	serviceToken, err := service.serviceToken.GetTokenByUserId(user.Id)
	if err != nil {
		return "", err
	}
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
		return "", schemas.ErrorAlreadyExistingRessource
	}
	workflowId, err := service.repository.SaveWorkflow(newWorkflow)
	if err != nil {
		return "", err
	}

	newWorkflow.Id = workflowId
	service.InitWorkflow(newWorkflow, serviceToken, result.ActionOption, result.ReactionOption)
	return "Workflow Created succesfully", nil

}

func (service *workflowService) ActivateWorkflow(ctx *gin.Context) error {
	result := schemas.WorkflowActivate{}
	err := ctx.ShouldBind(&result)
	if err != nil {
		return schemas.ErrorBadParameter
	}
	if !result.WorkflowState && result.WorkflowState {
		return schemas.ErrorBadParameter
	}

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return err
	}

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return schemas.ErrUserNotFound
	}
	workflow, err := service.repository.FindByIds(result.WorkflowId)
	if err != nil || workflow.Id == 0 {
		return schemas.ErrorNoWorkflowFound
	}
	newWorkflow := schemas.Workflow{
		Id:              workflow.Id,
		UserId:          user.Id,
		IsActive:        result.WorkflowState,
		ReactionTrigger: result.WorkflowState,
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
				reaction(channel, workflow.Id, githubServiceToken, reactionOption)
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
	result := ctx.Query("workflow_id")

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	workflowId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return nil, schemas.ErrorBadParameter
	}
	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return nil, schemas.ErrUserNotFound
	}
	workflow := service.repository.FindById(workflowId)
	if workflow.Id == 0 || workflow.UserId != user.Id {
		return nil, schemas.ErrorNoWorkflowFound
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
	err := ctx.ShouldBind(&result)
	if err != nil {
		return schemas.ErrorBadParameter
	}

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return err
	}

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return schemas.ErrUserNotFound
	}
	workflow := service.repository.FindById(result.WorkflowId)
	if workflow.Id == 0 {
		return schemas.ErrorNoWorkflowFound
	}
	if workflow.UserId == user.Id && workflow.ReactionId == result.ReactionId && workflow.ActionId == result.ActionId {
		actualReactionData := service.reactionResponseDataService.FindByWorkflowId(workflow.Id)
		for _, data := range actualReactionData {
			service.reactionResponseDataService.Delete(data)
		}
		err := service.repository.Delete(workflow.Id)
		if err != nil {
			return err
		}
		return nil
	}
	return schemas.ErrorNoWorkflowFound
}

func (service *workflowService) Update(ctx *gin.Context) error {
	result := schemas.WorkflowUpdateJson{}
	err := ctx.ShouldBind(&result)
	if err != nil {
		return schemas.ErrorBadParameter
	}

	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return err
	}

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return schemas.ErrUserNotFound
	}
	workflow := service.repository.FindById(result.WorkflowId)
	if workflow.Id == 0 {
		return schemas.ErrorNoWorkflowFound
	}
	if workflow.Id == result.WorkflowId && user.Id == workflow.UserId {
		workflow.ActionOptions = json.RawMessage(result.ActionOption)
		workflow.ReactionOptions = json.RawMessage(result.ReactionOption)
		workflow.Name = result.Name
		service.repository.Update(workflow)
		return nil
	}
	return schemas.ErrorNoWorkflowFound
}

func (service *workflowService) GetAllReactionsForAWorkflow(ctx *gin.Context) ([]json.RawMessage, error) {
	tokenString, err := toolbox.GetBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	user, err := service.userService.GetUserInfos(tokenString)
	if err != nil {
		return nil, schemas.ErrUserNotFound
	}
	workflow := service.repository.FindByUserId(user.Id)
	if len(workflow) == 0 {
		return []json.RawMessage{}, nil
	}
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
