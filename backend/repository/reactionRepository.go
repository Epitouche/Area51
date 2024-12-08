package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area51/schemas"
)


type ReactionRepository interface {
	Save(action schemas.Reaction)
	Update(action schemas.Reaction)
	Delete(action schemas.Reaction)
	FindAll() []schemas.Reaction
	FindByName(actionName string) []schemas.Reaction
	FindByServiceId(serviceId uint64) []schemas.Reaction
	FindByServiceByName(serviceID uint64, actionName string) []schemas.Reaction
	FindById(actionId uint64) schemas.Reaction
}

type reactionRepository struct {
	db *schemas.Database
}

func NewReactionRepository(conn *gorm.DB) ReactionRepository {
	err := conn.AutoMigrate(&schemas.Reaction{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &reactionRepository{
		db: &schemas.Database{
			Connection: conn,
		},
	}
}

func (repo *reactionRepository) Save(action schemas.Reaction) {
	err := repo.db.Connection.Create(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Update(action schemas.Reaction) {
	err := repo.db.Connection.Save(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Delete(action schemas.Reaction) {
	err := repo.db.Connection.Delete(&action)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) FindAll() []schemas.Reaction {
	var action []schemas.Reaction
	err := repo.db.Connection.Preload("Service").Find(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}

func (repo *reactionRepository) FindByName(actionName string) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{Name: actionName}).Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}

func (repo *reactionRepository) FindByServiceId(serviceId uint64) []schemas.Reaction {
	var reactions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{ServiceId: serviceId}).
		Find(&reactions)
	if err.Error != nil {
		panic(fmt.Errorf("failed to find action by service id: %v", err.Error))
	}
	return reactions
}

func (repo *reactionRepository) FindByServiceByName(
	serviceID uint64,
	actionName string,
) []schemas.Reaction {
	var actions []schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{ServiceId: serviceID, Name: actionName}).
		Find(&actions)
	if err.Error != nil {
		panic(err.Error)
	}
	return actions
}

func (repo *reactionRepository) FindById(actionId uint64) schemas.Reaction {
	var action schemas.Reaction
	err := repo.db.Connection.Where(&schemas.Reaction{Id: actionId}).First(&action)
	if err.Error != nil {
		panic(err.Error)
	}
	return action
}