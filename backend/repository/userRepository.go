package repository

import (
	"fmt"

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
	FindByEmail(email *string) schemas.User
	FindAllServicesByUserId(id uint64) []schemas.ServiceToken
	FindAllWorkflowsByUserId(id uint64) []schemas.Workflow
	AddServiceToUser(user schemas.User, service schemas.ServiceToken)
	GetAllServicesForUser(userId uint64) ([]schemas.ServiceToken, error)
	GetServiceByIdForUser(user schemas.User, serviceId uint64) (schemas.ServiceToken, error)
	LogoutFromService(user schemas.User, serviceToDelete schemas.Service) error
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
	err := r.db.Connection.Where(&schemas.User{
		Id: user.Id,
	}).Updates(&user)

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

func (r *userRepository) FindAll() (users []schemas.User) {
	err := r.db.Connection.Find(&users)

	if err.Error != nil {
		return []schemas.User{}
	}
	return users
}

func (r *userRepository) FindById(id uint64) (user schemas.User) {
	err := r.db.Connection.First(&user, id)

	if err.Error != nil {
		return schemas.User{}
	}
	return user
}

func (r *userRepository) FindByUsername(username string) (user schemas.User) {
	err := r.db.Connection.Where(&schemas.User{
		Username: username,
	}).First(&user)

	if err.Error != nil {
		return schemas.User{}
	}
	return user
}

func (r *userRepository) FindByEmail(email *string) (user schemas.User) {
	err := r.db.Connection.Where(&schemas.User{
		Email: email,
	}).First(&user)

	if err.Error != nil {
		return schemas.User{}
	}
	return user
}

func (r *userRepository) FindAllServicesByUserId(id uint64) []schemas.ServiceToken {
	user := r.FindById(id)
	err := r.db.Connection.Model(&user).Association("Services").Find(&user.Services)
	if err != nil {
		return []schemas.ServiceToken{}
	}
	for _, service := range user.Services {
		if service.UserId == id {
			return user.Services
		}
	}
	return []schemas.ServiceToken{}
}

func (r *userRepository) FindAllWorkflowsByUserId(id uint64) []schemas.Workflow {
	var workflows []schemas.Workflow
	err := r.db.Connection.Where(&schemas.Workflow{UserId: id}).Find(&workflows)
	if err.Error != nil || len(workflows) == 0 {
		return []schemas.Workflow{}
	}
	return workflows
}

func (r *userRepository) AddServiceToUser(user schemas.User, service schemas.ServiceToken) {
	r.db.Connection.Model(&user).Association("Services").Append(&service)
	r.db.Connection.Save(&user)
}

func (r *userRepository) GetAllServicesForUser(userId uint64) ([]schemas.ServiceToken, error) {
	var services []schemas.ServiceToken
	err := r.db.Connection.Where(&schemas.ServiceToken{UserId: userId}).Find(&services)
	if err.Error != nil {
		return []schemas.ServiceToken{}, err.Error
	}
	return services, nil
}

func (r *userRepository) GetServiceByIdForUser(user schemas.User, serviceId uint64) (schemas.ServiceToken, error) {
	err := r.db.Connection.Model(&user).Association("Services").Find(&user.Services)
	if err != nil {
		return schemas.ServiceToken{}, err
	}
	for _, service := range user.Services {
		if service.ServiceId == serviceId {
			return service, nil
		}
	}
	return schemas.ServiceToken{}, nil
}

func (r *userRepository) LogoutFromService(user schemas.User, serviceToDelete schemas.Service) error {
	r.db.Connection.Model(&user).Association("Services").Find(&user.Services)

	for _, service := range user.Services {
		if service.ServiceId == serviceToDelete.Id {
			r.db.Connection.Model(&user).Association("Services").Delete(&service)
			return nil
		}
	}
	return fmt.Errorf("service not found")
}
