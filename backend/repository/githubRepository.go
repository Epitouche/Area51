package repository

import (
	"area51/schemas"

	"gorm.io/gorm"
)

type GithubTokenRepository interface {
	Save(token schemas.OAuth2Token)
	FindByAccessToken(accessToken string) []schemas.OAuth2Token
}

type githubTokenRepositoryStruct struct {
	db *schemas.Database
}

func NewGithubTokenRepository(conn *gorm.DB) GithubTokenRepository {
	err := conn.AutoMigrate(&schemas.OAuth2Token{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &githubTokenRepositoryStruct{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *githubTokenRepositoryStruct) Save(token schemas.OAuth2Token) {
	err := repo.db.Connection.Create(&token)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *githubTokenRepositoryStruct) FindByAccessToken(accessToken string) []schemas.OAuth2Token {
	var tokens []schemas.OAuth2Token
	err := repo.db.Connection.Where(&schemas.OAuth2Token{AccessToken: accessToken}).Find(&tokens)
	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}