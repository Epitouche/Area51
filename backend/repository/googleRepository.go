package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type GoogleRepository interface {
	Save(googleMailOptions schemas.GoogleActionResponse)
	Update(googleMailOptions schemas.GoogleActionResponse)
	Delete(googleMailOptions schemas.GoogleActionResponse)
	UpdateNumEmails(googleMailOptions schemas.GoogleActionResponse)

	FindAll() []schemas.GoogleActionResponse
	FindByWorkflowId(workflowId uint64) schemas.GoogleActionResponse
}

type googleRepository struct {
	db *schemas.Database
}

func NewGoogleRepository(db *gorm.DB) GoogleRepository {
	err := db.AutoMigrate(&schemas.GoogleActionResponse{})
	if err != nil {
		panic("failed to migrate database")
	}

	return &googleRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}

func (repo *googleRepository) Save(googleMailOptions schemas.GoogleActionResponse) {
	err := repo.db.Connection.Create(&googleMailOptions)

	if err.Error != nil {
		return
	}
}

func (repo *googleRepository) Update(googleMailOptions schemas.GoogleActionResponse) {
	err := repo.db.Connection.Where(&schemas.GoogleActionResponse{
		Id: googleMailOptions.Id,
	}).Updates(&googleMailOptions)

	if err.Error != nil {
		return
	}
}

func (repo *googleRepository) Delete(googleMailOptions schemas.GoogleActionResponse) {
	err := repo.db.Connection.Delete(&googleMailOptions)

	if err.Error != nil {
		return
	}
}

func (repo *googleRepository) UpdateNumEmails(googleMailOptions schemas.GoogleActionResponse) {
	err := repo.db.Connection.Model(&schemas.GoogleActionResponse{}).Where(&schemas.GoogleActionResponse{Id: googleMailOptions.Id}).Updates(map[string]interface{}{
		"result_size_estimate": googleMailOptions.ResultSizeEstimate,
	})
	if err.Error != nil {
		return
	}
}

func (repo *googleRepository) FindAll() (googleMailOptions []schemas.GoogleActionResponse) {
	err := repo.db.Connection.Find(&googleMailOptions)

	if err.Error != nil {
		return []schemas.GoogleActionResponse{}
	}
	return googleMailOptions
}

func (repo *googleRepository) FindByWorkflowId(workflowId uint64) (googleMailOptions schemas.GoogleActionResponse) {
	err := repo.db.Connection.Where(&schemas.GoogleActionResponse{WorkflowId: workflowId}).First(&googleMailOptions)

	if err.Error != nil {
		return schemas.GoogleActionResponse{}
	}
	return googleMailOptions
}
