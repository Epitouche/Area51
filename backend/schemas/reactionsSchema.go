package schemas

import (
	"encoding/json"
	"time"
)

type ReactionJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ReactionResponseData struct {
	Id 				uint64 			`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	WorkflowId 		uint64 			`json:"workflow_id"`
	Workflow 		Workflow 		`json:"workflow,omitempty" gorm:"foreignkey:WorkflowId;references:Id"`
	ApiResponse  	json.RawMessage `gorm:"type:jsonb" json:"apiResponse"`
	CreatedAt		time.Time       `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type Reaction struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	ServiceId		uint64		`json:"-"`
	Service			Service		`json:"service,omitempty" gorm:"foreignkey:ServiceId;references:Id"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
	Description		string		`json:"description" gorm:"type:varchar(100)"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Trigger			string		`json:"trigger"`
}
