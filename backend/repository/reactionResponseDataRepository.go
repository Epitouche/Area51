package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type ReactionResponseDataRepository interface {
	Save(reactionResponseData schemas.ReactionResponseData)
	Update(reactionResponseData schemas.ReactionResponseData)
	Delete(reactionResponseData schemas.ReactionResponseData)
	FindAll() []schemas.ReactionResponseData
	FindByWorkflowId(workflowId uint64) []schemas.ReactionResponseData
}

type reactionResponseDataRepository struct {
	db *schemas.Database
}

func NewReactionResponseDataRepository(conn *gorm.DB) ReactionResponseDataRepository {
	err := conn.AutoMigrate(&schemas.ReactionResponseData{})

	if err != nil {
		panic("failed to migrate database")
	}

	return &reactionResponseDataRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *reactionResponseDataRepository) Save(reactionResponseData schemas.ReactionResponseData) {
	err := repo.db.Connection.Create(&reactionResponseData)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionResponseDataRepository) Update(reactionResponseData schemas.ReactionResponseData) {
	err := repo.db.Connection.Save(&reactionResponseData)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionResponseDataRepository) Delete(reactionResponseData schemas.ReactionResponseData) {
	err := repo.db.Connection.Delete(&reactionResponseData)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionResponseDataRepository) FindAll() (reactionResponseData []schemas.ReactionResponseData) {
	err := repo.db.Connection.Find(&reactionResponseData)

	if err.Error != nil {
		panic(err.Error)
	}
	return reactionResponseData
}

func (repo *reactionResponseDataRepository) FindByWorkflowId(workflowId uint64) (reactionResponseData []schemas.ReactionResponseData) {
	err := repo.db.Connection.Where(&schemas.ReactionResponseData{
		WorkflowId: workflowId,
	}).Find(&reactionResponseData)

	if err.Error != nil {
		panic(err.Error)
	}
	return reactionResponseData
}
