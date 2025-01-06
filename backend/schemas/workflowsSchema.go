package schemas

import "time"

type WorkflowResult struct {
	UserId         uint64 `json:"user_id"`
	ActionOption   string `json:"action_option"   binding:"required"`
	ActionId       uint64 `json:"action_id"`
	ReactionOption string `json:"reaction_option" binding:"required"`
	Name           string `json:"name"`
	ReactionId     uint64 `json:"reaction_id"`
}

type WorkflowJson struct {
	Name       string    `json:"name"`
	ActionId   uint64    `json:"action_id"`
	ReactionId uint64    `json:"reaction_id"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
}

type Workflow struct {
	Id         uint64    `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId     uint64    `json:"-"`
	User       User      `json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	ActionId   uint64    `json:"-"`
	Action     Action    `json:"action,omitempty" gorm:"foreignkey:ActionId;references:Id"`
	ReactionId uint64    `json:"-"`
	Reaction   Reaction  `json:"reaction,omitempty" gorm:"foreignkey:ReactionId;references:Id"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive   bool      `json:"is_active" default:"false" gorm:"column:is_active"`
	Name       string    `json:"name" gorm:"type:varchar(100)"`
}
