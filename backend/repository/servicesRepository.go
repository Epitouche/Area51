package repository

import (
	"gorm.io/gorm"

	"area51/schemas"
)

type ServiceRepository interface {
	Save(service schemas.Service)
	Update(service schemas.Service)
	Delete(service schemas.Service)
	FindAll() []schemas.Service
	FindByName(serviceName schemas.ServiceName) schemas.Service
	FindAllByName(serviceName schemas.ServiceName) []schemas.Service
	FindById(serviceId uint64) schemas.Service
}

type serviceRepository struct {
	db *schemas.Database
}


func NewServiceRepository(db *gorm.DB) ServiceRepository {
	err := db.AutoMigrate(&schemas.Service{})
	if err != nil {
		panic(err)
	}
	return &serviceRepository{
		db: &schemas.Database{
			Connection: db,
		},
	}
}

func (repo *serviceRepository) Save(service schemas.Service) {
	err := repo.db.Connection.Create(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) Update(service schemas.Service) {
	err := repo.db.Connection.Save(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) Delete(service schemas.Service) {
	err := repo.db.Connection.Delete(&service)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *serviceRepository) FindAll() []schemas.Service {
	var services []schemas.Service
	err := repo.db.Connection.Find(&services)
	if err.Error != nil {
		panic(err.Error)
	}
	return services
}

func (repo *serviceRepository) FindByName(serviceName schemas.ServiceName) schemas.Service {
	var service schemas.Service
	err := repo.db.Connection.Where(&schemas.Service{Name: serviceName}).First(&service)
	if err.Error != nil {
		panic(err.Error)
	}
	return service
}

func (repo *serviceRepository) FindAllByName(serviceName schemas.ServiceName) []schemas.Service {
	var services []schemas.Service
	err := repo.db.Connection.Where(&schemas.Service{Name: serviceName}).Find(&services)
	if err.Error != nil {
		panic(err.Error)
	}
	return services
}

func (repo *serviceRepository) FindById(serviceId uint64) schemas.Service {
	var service schemas.Service
	err := repo.db.Connection.Where(&schemas.Service{Id: serviceId}).First(&service)
	if err.Error != nil {
		panic(err.Error)
	}
	return service
}
