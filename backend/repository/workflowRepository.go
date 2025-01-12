package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area51/schemas"
)

type WorkflowRepository interface {
	Save(workflow schemas.Workflow)
	Update(workflow schemas.Workflow)
	UpdateActiveStatus(workflow schemas.Workflow)
	Delete(workflowId uint64) error

	FindAll() []schemas.Workflow
	FindByIds(workflowId uint64) (schemas.Workflow, error)
	FindByUserId(userId uint64) []schemas.Workflow
	FindByWorkflowName(workflowName string) schemas.Workflow
	FindById(workflowId uint64) schemas.Workflow
	FindByActionId(actionId uint64) schemas.Workflow
	FindByReactionId(reactionId uint64) schemas.Workflow
	SaveWorkflow(workflow schemas.Workflow) (workflowId uint64, err error)
	FindExistingWorkflow(workflow schemas.Workflow) schemas.Workflow
}

type workflowRepository struct {
	db *schemas.Database
}

func NewWorkflowRepository(db *gorm.DB) WorkflowRepository {
	err := db.AutoMigrate(&schemas.Workflow{})

	if err != nil {
		panic("failed to migrate database")
	}

	return &workflowRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}

func (repo *workflowRepository) Save(workflow schemas.Workflow) {
	err := repo.db.Connection.Create(&workflow)

	if err.Error != nil {
		fmt.Printf("%+v", err.Error)
		panic(err.Error)
	}
}

func (repo *workflowRepository) Update(workflow schemas.Workflow) {
	err := repo.db.Connection.Where(&schemas.Workflow{
		Id: workflow.Id,
	}).Updates(&workflow)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *workflowRepository) UpdateActiveStatus(workflow schemas.Workflow) {
	err := repo.db.Connection.Model(&schemas.Workflow{}).Where(&schemas.Workflow{Id: workflow.Id}).Updates(map[string]interface{}{
		"is_active": workflow.IsActive,
	})
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *workflowRepository) Delete(workflowId uint64) error {
	err := repo.db.Connection.Delete(&schemas.Workflow{
		Id: workflowId,
	})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *workflowRepository) FindAll() (workflows []schemas.Workflow) {
	err := repo.db.Connection.Find(&workflows)

	if err.Error != nil {
		return []schemas.Workflow{}
	}
	return workflows
}

func (repo *workflowRepository) FindByIds(workflowId uint64) (schemas.Workflow, error) {
	workflow := schemas.Workflow{}

	err := repo.db.Connection.Where(&schemas.Workflow{Id: workflowId}).First(&workflow)
	if err.Error != nil {
		return schemas.Workflow{}, err.Error
	}

	err = repo.db.Connection.Where(&schemas.Action{Id: workflow.ActionId}).First(&workflow.Action)
	if err.Error != nil {
		return schemas.Workflow{}, err.Error
	}

	err = repo.db.Connection.Where(&schemas.Reaction{Id: workflow.ReactionId}).First(&workflow.Reaction)
	if err.Error != nil {
		return schemas.Workflow{}, err.Error
	}

	return workflow, nil
}

func (repo *workflowRepository) FindExistingWorkflow(workflow schemas.Workflow) schemas.Workflow {
	err := repo.db.Connection.Where(&schemas.Workflow{UserId: workflow.UserId, ActionId: workflow.ActionId, ReactionId: workflow.ReactionId}).First(&workflow)
	if err.Error != nil {
		return schemas.Workflow{}
	}
	return workflow
}

func (repo *workflowRepository) FindByUserId(userId uint64) []schemas.Workflow {
	var workflows []schemas.Workflow
	err := repo.db.Connection.Where(&schemas.Workflow{UserId: userId}).Find(&workflows)
	if err.Error != nil {
		return []schemas.Workflow{}
	}
	return workflows
}

func (repo *workflowRepository) FindByWorkflowName(workflowName string) (workflow schemas.Workflow) {
	err := repo.db.Connection.Where(&schemas.Workflow{
		Name: workflowName,
	}).First(&workflow)

	if err.Error != nil {
		return schemas.Workflow{}
	}
	return workflow
}

func (repo *workflowRepository) FindById(workflowId uint64) schemas.Workflow {
	var workflow schemas.Workflow
	err := repo.db.Connection.Where(&schemas.Workflow{Id: workflowId}).First(&workflow)
	if err.Error != nil {
		return schemas.Workflow{}
	}
	return workflow
}

func (repo *workflowRepository) FindByActionId(actionId uint64) schemas.Workflow {
	var workflow schemas.Workflow
	err := repo.db.Connection.Where(&schemas.Workflow{ActionId: actionId}).First(&workflow)
	if err.Error != nil {
		return schemas.Workflow{}
	}
	return workflow
}

func (repo *workflowRepository) FindByReactionId(reactionId uint64) (workflow schemas.Workflow) {
	err := repo.db.Connection.Where(&schemas.Workflow{
		ReactionId: reactionId,
	}).First(&workflow)

	if err.Error != nil {
		return schemas.Workflow{}
	}
	return workflow
}

func (repo *workflowRepository) SaveWorkflow(workflow schemas.Workflow) (workflowId uint64, err error) {
	repo.Save(workflow)
	result := repo.db.Connection.Last(&workflow)

	if result.Error != nil {
		return 0, result.Error
	}
	return workflow.Id, nil
}
