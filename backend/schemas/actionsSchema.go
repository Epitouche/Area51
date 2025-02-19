package schemas

import (
	"encoding/json"
	"errors"
	"time"
)

type ActionJson struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Options     json.RawMessage `json:"options"`
	ActionId    uint64          `json:"action_id"`
}

type ActionResult struct {
	UserId      uint64 `json:"-"`
	ServiceName string `json:"serviceName"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Options     string `json:"options"`
}

type Action struct {
	Id          uint64          `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name        string          `json:"name" gorm:"type:varchar(100)"`
	ServiceId   uint64          `json:"-"`
	Service     Service         `json:"service,omitempty" gorm:"foreignkey:ServiceId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
	Description string          `json:"description" gorm:"type:varchar(100)"`
	CreatedAt   time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Options     json.RawMessage `gorm:"type:jsonb" json:"options"`
}

var (
	ErrActionNotFound = errors.New("action not found")
)
