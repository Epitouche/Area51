package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type TokenRepository interface {
	Save(token schemas.ServiceToken)
	Update(token schemas.ServiceToken)
	Delete(token schemas.ServiceToken)
	FindAll() []schemas.ServiceToken
	FindByToken(token string) []schemas.ServiceToken
	FindById(tokenId uint64) schemas.ServiceToken
	FindByUserId(user schemas.User) []schemas.ServiceToken
	FindByUserIdAndServiceId(userId uint64, serviceId uint64) schemas.ServiceToken
}

type tokenRepository struct {
	db *schemas.Database
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	err := db.AutoMigrate(&schemas.ServiceToken{})

	if err != nil {
		panic(err)
	}

	return &tokenRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}

func (repo *tokenRepository) Save(token schemas.ServiceToken) {
	err := repo.db.Connection.Create(&token)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) Update(token schemas.ServiceToken) {
	err := repo.db.Connection.Where(&schemas.ServiceToken{
		Id: token.Id,
	}).Updates(&token)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) Delete(token schemas.ServiceToken) {
	err := repo.db.Connection.Delete(&token)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *tokenRepository) FindAll() (tokens []schemas.ServiceToken) {
	err := repo.db.Connection.Find(&tokens)

	if err.Error != nil {
		panic(err.Error)
	}
	return tokens
}

func (repo *tokenRepository) FindByToken(token string) (serviceTokens []schemas.ServiceToken) {
	err := repo.db.Connection.Where(&schemas.ServiceToken{
		Token: token,
	}).Find(&serviceTokens)

	if err.Error != nil {
		return []schemas.ServiceToken{}
	}
	return serviceTokens
}

func (repo *tokenRepository) FindById(tokenId uint64) (serviceToken schemas.ServiceToken) {
	err := repo.db.Connection.First(&serviceToken, tokenId)

	if err.Error != nil {
		return schemas.ServiceToken{}
	}
	return serviceToken
}

func (repo *tokenRepository) FindByUserId(user schemas.User) (serviceTokens []schemas.ServiceToken) {
	var services []schemas.ServiceToken
	err := repo.db.Connection.Model(&user).Association("Services").Find(&services)
	if err != nil {
		return []schemas.ServiceToken{}
	}
	return services
}

func (repo *tokenRepository) FindByUserIdAndServiceId(userId uint64, serviceId uint64) (serviceToken schemas.ServiceToken) {
	err := repo.db.Connection.Where(&schemas.ServiceToken{
		UserId:    userId,
		ServiceId: serviceId,
	}).First(&serviceToken)

	if err.Error != nil {
		return schemas.ServiceToken{}
	}
	return serviceToken
}
