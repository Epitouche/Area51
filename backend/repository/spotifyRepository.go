package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type SpotifyRepository interface {
}

type spotifyRepository struct {
	db *schemas.Database
}

func NewSpotifyRepository(db *gorm.DB) SpotifyRepository {
	return &spotifyRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}
