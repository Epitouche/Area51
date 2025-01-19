package schemas

import (
	"encoding/json"
	"errors"
	"time"
)

type WorkflowResult struct {
	ActionOption   json.RawMessage `gorm:"type:jsonb" json:"action_option"   binding:"required"`
	ActionId       uint64          `json:"action_id" binding:"required"`
	ReactionOption json.RawMessage `gorm:"type:jsonb" json:"reaction_option" binding:"required"`
	Name           string          `json:"name"`
	ReactionId     uint64          `json:"reaction_id" binding:"required"`
}

type WorkflowActivate struct {
	WorkflowId    uint64 `json:"workflow_id" binding:"required"`
	WorkflowState bool   `json:"workflow_state" binding:"required"`
}

type WorkflowJson struct {
	Name           string          `json:"name"`
	WorkflowId     uint64          `json:"workflow_id" binding:"required"`
	ActionName     string          `json:"action_name"`
	ActionId       uint64          `json:"action_id" binding:"required"`
	ActionOption   json.RawMessage `json:"action_option"`
	ReactionId     uint64          `json:"reaction_id" binding:"required"`
	ReactionName   string          `json:"reaction_name"`
	ReactionOption json.RawMessage `json:"reaction_option"`
	IsActive       bool            `json:"is_active"`
	CreatedAt      time.Time       `json:"created_at"`
}

type WorkflowUpdateJson struct {
	WorkflowId     uint64          `json:"workflow_id" binding:"required"`
	ActionOption   json.RawMessage `json:"action_option" binding:"required"`
	ReactionOption json.RawMessage `json:"reaction_option" binding:"required"`
	Name           string          `json:"name" binding:"required"`
}

type Workflow struct {
	Id              uint64          `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId          uint64          `json:"-"`
	User            User            `json:"user,omitempty" gorm:"foreignkey:UserId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	ActionId        uint64          `json:"-"`
	Action          Action          `json:"action,omitempty" gorm:"foreignkey:ActionId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	ActionOptions   json.RawMessage `gorm:"type:jsonb" json:"action_options"`
	ReactionId      uint64          `json:"-"`
	Reaction        Reaction        `json:"reaction,omitempty" gorm:"foreignkey:ReactionId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	ReactionOptions json.RawMessage `gorm:"type:jsonb" json:"reaction_options"`
	CreatedAt       time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive        bool            `json:"is_active" default:"false" gorm:"column:is_active"`
	ReactionTrigger bool            `json:"reaction_trigger" default:"false" gorm:"column:reaction_trigger"`
	Name            string          `json:"name" gorm:"type:varchar(100)"`
	Utils           json.RawMessage `gorm:"type:jsonb" json:"utils"`
}

var (
	ErrorBadParameter             = errors.New("invalid JSON format or structure")
	ErrorNoWorkflowFound          = errors.New("no workflow found")
	ErrorAlreadyExistingRessource = errors.New("ressource already exists")
)
