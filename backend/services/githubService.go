package services

import (
	"area51/repository"
	"area51/schemas"
	"errors"
)

type GithubTokenService interface {
	SaveToken(schemas.OAuth2Token) (tokenId uint64, err error)
}

type githubTokenServiceStruct struct {
	repository repository.GithubTokenRepository
}

func NewGithubTokenService(repo repository.GithubTokenRepository) GithubTokenService {
	return &githubTokenServiceStruct{
		repository: repo,
	}
}

func (service *githubTokenServiceStruct) SaveToken(token schemas.OAuth2Token) (tokenId uint64, err error) {
	tokens := service.repository.FindByAccessToken(token.AccessToken)

	for _, t := range tokens {
		if t.AccessToken == token.AccessToken {
			return t.Id, errors.New("token already exists")
		}
	}
	service.repository.Save(token)
	tokens = service.repository.FindByAccessToken(token.AccessToken)
	for _, t := range tokens {
		if t.AccessToken == token.AccessToken {
			return t.Id, nil
		}
	}
	return 0, errors.New("unable to save token")
}
