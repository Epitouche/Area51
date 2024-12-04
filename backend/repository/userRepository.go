package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type UserRepository interface {
	Save(user schemas.User)
	Update(user schemas.User)
	Delete(user schemas.User)

	FindAll() []schemas.User
	FindByID(id uint64) schemas.User
	FindByUsername(username string) schemas.User
	FindByEmail(email string) schemas.User
}

type userRepository struct {
	db *schemas.Database
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	err := conn.AutoMigrate(&schemas.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &userRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (r *userRepository) Save(user schemas.User) {
	err := r.db.Connection.Create(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (r *userRepository) Update(user schemas.User) {
	err := r.db.Connection.Save(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (r *userRepository) Delete(user schemas.User) {
	err := r.db.Connection.Delete(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (r *userRepository) FindAll() []schemas.User {
	var users []schemas.User
	err := r.db.Connection.Find(&users)
	if err.Error != nil {
		return []schemas.User{}
	}
	return users
}

func (r *userRepository) FindByID(id uint64) schemas.User {
	var user schemas.User
	err := r.db.Connection.First(&user, id)
	if err.Error != nil {
		return schemas.User{}
	}
	return user
}

func (r *userRepository) FindByUsername(username string) schemas.User {
	var user schemas.User
	err := r.db.Connection.Where(&schemas.User{Username: username}).First(&user)
	if err.Error != nil {
		return schemas.User{}
	}
	return user
}

func (r *userRepository) FindByEmail(email string) schemas.User {
	var user schemas.User
	err := r.db.Connection.Where(&schemas.User{Email: email}).First(&user)
	if err.Error != nil {
		return schemas.User{}
	}
	return user
}
