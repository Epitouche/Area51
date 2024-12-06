package schemas

import "time"

type WorkflowResult struct {
	UserId         uint64 `json:"-"`
	ActionOption   string `json:"action_option"   binding:"required"`
	ActionId       uint64 `json:"-"`
	ReactionOption string `json:"reaction_option" binding:"required"`
	ReactionId     uint64 `json:"-"`
}

type Workflow struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId			uint64		`json:"-"`
	User			User		`json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	ActionId		uint64		`json:"-"`
	Action			Action		`json:"action,omitempty" gorm:"foreignkey:ActionId;references:Id"`
	ReactionId		uint64		`json:"-"`
	Reaction		Reaction	`json:"reaction,omitempty" gorm:"foreignkey:ReactionId;references:Id"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive		bool		`json:"is_active" gorm:"type:boolean"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
}
