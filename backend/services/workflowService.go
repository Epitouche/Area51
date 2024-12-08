package services

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"area51/repository"
	"area51/schemas"
)

type WorkflowService interface {
	FindAll() []schemas.Workflow
	CreateWorkflow(ctx *gin.Context) (string, error)
	InitWorkflow(workflowStartingPoint schemas.Workflow)
	ExistWorkflow(workflowId uint64) bool
}

type workflowService struct {
	repository 		repository.WorkflowRepository
	userService 	UserService
	actionService 	ActionService
	reactionService ReactionService
	servicesService ServicesService
}

func NewWorkflowService(
	repository repository.WorkflowRepository,
	userService UserService,
	actionService ActionService,
	reactionService ReactionService,
	servicesService ServicesService,
	) WorkflowService {
	return &workflowService{
		repository: repository,
		userService: userService,
		actionService: actionService,
		reactionService: reactionService,
		servicesService: servicesService,
	}
}

func (service *workflowService) FindAll() []schemas.Workflow {
	return service.repository.FindAll()
}

func (service *workflowService) CreateWorkflow(ctx *gin.Context) (string, error) {
	var result schemas.WorkflowResult
	err := json.NewDecoder(ctx.Request.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println(result)
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	user, err:= service.userService.GetUserInfos(tokenString)
	if err != nil {
		return "", err
	}

	newWorkflow := schemas.Workflow{
		UserId: result.UserId,
		User: user,
		IsActive: true,
		ActionId: result.ActionId,
		ReactionId: result.ReactionId,
		Action: service.actionService.FindById(result.ActionId),
		Reaction: service.reactionService.FindById(result.ReactionId),

	}
	workflowId, err := service.repository.SaveWorkflow(newWorkflow)
	if err != nil {
		return "", err
	}
	newWorkflow.Id = workflowId
	service.InitWorkflow(newWorkflow)
	return "Workflow Created succesfully", nil

}

func (service *workflowService) InitWorkflow(workflowStartingPoint schemas.Workflow) {
	workflowChannel := make(chan string)
	go service.WorkflowActionChannel(workflowStartingPoint, workflowChannel)
	go service.WorkflowReactionChannel(workflowStartingPoint, workflowChannel)
}

func (service *workflowService) WorkflowActionChannel(workflowStartingPoint schemas.Workflow, channel chan string) {
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
                fmt.Println("Action not found")
                return
            }
            if workflow.IsActive {
                action(channel, workflow.Action.Name, workflow.Id)
            }
        }
        fmt.Println("Clear")
        channel <- "Workflow finished"
    }(workflowStartingPoint, channel)
}

func (service *workflowService) WorkflowReactionChannel(workflowStartingPoint schemas.Workflow, channel chan string) {
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
				reaction(workflow.Id)
			}
		}
		fmt.Println("Clear")
		channel <- "Workflow finished"
	}(workflowStartingPoint, channel)
}

func (service *workflowService) ExistWorkflow(workflowId uint64) bool {
	_, err := service.repository.FindByIds(workflowId)
	return err == nil
}