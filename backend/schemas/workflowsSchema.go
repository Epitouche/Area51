package schemas

import (
	"encoding/json"
	"time"
)

type WorkflowResult struct {
	UserId         uint64          `json:"user_id"`
	ActionOption   json.RawMessage `gorm:"type:jsonb" json:"action_option"   binding:"required"`
	ActionId       uint64          `json:"action_id"`
	ReactionOption json.RawMessage `gorm:"type:jsonb" json:"reaction_option" binding:"required"`
	Name           string          `json:"name"`
	ReactionId     uint64          `json:"reaction_id"`
}

type WorkflowActivate struct {
	WorkflowId   uint64 `json:"workflow_id" binding:"required"`
	WorflowState bool   `json:"workflow_state" binding:"required"`
}

type WorkflowJson struct {
	Name         string    `json:"name"`
	WorkflowId   uint64    `json:"workflow_id"`
	ActionName   string    `json:"action_name"`
	ActionId     uint64    `json:"action_id"`
	ReactionId   uint64    `json:"reaction_id"`
	ReactionName string    `json:"reaction_name"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

type Workflow struct {
	Id              uint64          `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId          uint64          `json:"-"`
	User            User            `json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	ActionId        uint64          `json:"-"`
	Action          Action          `json:"action,omitempty" gorm:"foreignkey:ActionId;references:Id"`
	ActionOptions   json.RawMessage `gorm:"type:jsonb" json:"action_options"`
	ReactionId      uint64          `json:"-"`
	Reaction        Reaction        `json:"reaction,omitempty" gorm:"foreignkey:ReactionId;references:Id"`
	ReactionOptions json.RawMessage `gorm:"type:jsonb" json:"reaction_options"`
	CreatedAt       time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive        bool            `json:"is_active" default:"false" gorm:"column:is_active"`
	ReactionTrigger bool            `json:"reaction_trigger" default:"false" gorm:"column:reaction_trigger"`
	Name            string          `json:"name" gorm:"type:varchar(100)"`
}
