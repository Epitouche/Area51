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
	FindById(id uint64) schemas.User
	FindByUsername(username string) schemas.User
	FindByEmail(email string) schemas.User
	FindAllServicesByUserId(id uint64) []schemas.ServiceToken
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

func (repo *userRepository) Save(user schemas.User) {
	err := repo.db.Connection.Create(&user)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (r *userRepository) Update(user schemas.User) {
	err := r.db.Connection.Where(&schemas.User{Id: user.Id}).Updates(&user)
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

func (r *userRepository) FindById(id uint64) schemas.User {
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

func (r *userRepository) FindAllServicesByUserId(id uint64) []schemas.ServiceToken {
	var services []schemas.ServiceToken
	err := r.db.Connection.Where(&schemas.ServiceToken{UserId: id}).Find(&services)
	if err.Error != nil {
		return []schemas.ServiceToken{}
	}
	return services
}
