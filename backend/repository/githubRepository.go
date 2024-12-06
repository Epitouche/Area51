package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type GithubRepository interface {
}

type githubRepository struct {
	db *schemas.Database
}

func NewGithubRepository(db *gorm.DB) GithubRepository {
	return &githubRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}
