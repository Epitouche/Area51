package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type UserRepository interface {
	Save(user schemas.User)
	FindByEmail(email string) []schemas.User
	FindByUsername(username string) []schemas.User
}

type userRepositoryStruct struct {
	db *schemas.Database
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	err := connection.AutoMigrate(&schemas.User{})
	if err != nil {
		panic("Failed to migrate the database")
	}
	return &userRepositoryStruct{
		db: &schemas.Database{
			Connection: connection,
		},
	}
}

func (repo *userRepositoryStruct) Save(user schemas.User) {
	err := repo.db.Connection.Create(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *userRepositoryStruct) FindByEmail(email string) (users []schemas.User) {
	err := repo.db.Connection.Where(&schemas.User{Email: email}).Find(&users)
	if err.Error != nil {
		panic(err.Error)
	}
	return users
}

func (repo *userRepositoryStruct) FindByUsername(username string) (users []schemas.User) {
	err := repo.db.Connection.Where(&schemas.User{Username: username}).Find(&users)
	if err.Error != nil {
		panic(err.Error)
	}
	return users
}