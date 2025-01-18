package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type GoogleRepository interface {
}

type googleRepository struct {
	db *schemas.Database
}

func NewGoogleRepository(db *gorm.DB) GoogleRepository {

	return &googleRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}
