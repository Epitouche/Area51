package schemas

import "time"

type Reaction struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	ServiceId		uint64		`json:"service_id,omitempty"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
	Description		string		`json:"description" gorm:"type:varchar(100)"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Trigger			string		`json:"trigger"`

	Service			Service		`json:"service,omitempty" gorm:"foreignkey:ServiceId;references:Id"`
}
