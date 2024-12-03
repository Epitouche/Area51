package schemas

import "time"

type Workflow struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId			User		`json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	ActionID		Action		`json:"action,omitempty" gorm:"foreignkey:ActionId;references:Id"`
	ReactionId		Reaction	`json:"reaction,omitempty" gorm:"foreignkey:ReactionId;references:Id"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive		bool		`json:"is_active" gorm:"type:boolean"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
}
