package schemas

import "time"

type Workflow struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId			uint64		`json:"user_id,omitempty"`
	ActionId		uint64		`json:"action_id,omitempty"`
	ReactionId		uint64		`json:"reaction_id,omitempty"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive		bool		`json:"is_active" gorm:"type:boolean"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`

	User			User		`json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	Action			Action		`json:"action,omitempty" gorm:"foreignkey:ActionId;references:Id"`
	Reaction		Reaction	`json:"reaction,omitempty" gorm:"foreignkey:ReactionId;references:Id"`
}
