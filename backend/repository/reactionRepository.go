package repository

import (
	"fmt"

	"gorm.io/gorm"

	"area51/schemas"
)

type ReactionRepository interface {
	Save(action schemas.Reaction)
	Update(reaction schemas.Reaction)
	UpdateTrigger(reaction schemas.Reaction)
	Delete(action schemas.Reaction)
	FindAll() []schemas.Reaction
	FindByName(reactionName string) []schemas.Reaction
	FindAllByName(reactionName string) []schemas.Reaction
	FindByServiceId(serviceId uint64) []schemas.Reaction
	FindByServiceByName(serviceID uint64, reactionName string) []schemas.Reaction
	FindById(reactionId uint64) schemas.Reaction
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

func (repo *reactionRepository) Save(reaction schemas.Reaction) {
	err := repo.db.Connection.Create(&reaction)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Update(reaction schemas.Reaction) {
	err := repo.db.Connection.Where(&schemas.Reaction{Id: reaction.Id}).Updates(&reaction)
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) UpdateTrigger(reaction schemas.Reaction) {
	err := repo.db.Connection.Model(&schemas.Reaction{}).Where(&schemas.Reaction{Id: reaction.Id}).Updates(map[string]interface{}{
		"trigger": reaction.Trigger,
	})
	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) Delete(reaction schemas.Reaction) {
	err := repo.db.Connection.Delete(&reaction)

	if err.Error != nil {
		panic(err.Error)
	}
}

func (repo *reactionRepository) FindAll() (reactions []schemas.Reaction) {
	err := repo.db.Connection.Preload("Service").Find(&reactions)

	if err.Error != nil {
		panic(err.Error)
	}
	return reactions
}

func (repo *reactionRepository) FindByName(reactionName string) (reactions []schemas.Reaction) {
	err := repo.db.Connection.Where(&schemas.Reaction{
		Name: reactionName,
	}).Find(&reactions)

	if err.Error != nil {
		panic(err.Error)
	}
	return reactions
}

func (repo *reactionRepository) FindByServiceId(serviceId uint64) (reactions []schemas.Reaction) {
	err := repo.db.Connection.Where(&schemas.Reaction{
		ServiceId: serviceId,
	}).Find(&reactions)

	if err.Error != nil {
		panic(fmt.Errorf("failed to find reaction by service id: %v", err.Error))
	}
	return reactions
}

func (repo *reactionRepository) FindByServiceByName(
	serviceID uint64,
	reactionName string,
) (reactions []schemas.Reaction) {
	err := repo.db.Connection.Where(&schemas.Reaction{
		ServiceId: serviceID,
		Name: reactionName,
	}).Find(&reactions)

	if err.Error != nil {
		panic(err.Error)
	}
	return reactions
}

func (repo *reactionRepository) FindById(reactionId uint64) (reaction schemas.Reaction) {
	err := repo.db.Connection.Where(&schemas.Reaction{
		Id: reactionId,
	}).First(&reaction)

	if err.Error != nil {
		panic(err.Error)
	}
	return reaction
}

func (repo *reactionRepository) FindAllByName(reactionName string) (reactions []schemas.Reaction) {
	err := repo.db.Connection.Where(&schemas.Reaction{
		Name: reactionName,
	}).Find(&reactions)

	if err.Error != nil {
		panic(err.Error)
	}
	return reactions
}
