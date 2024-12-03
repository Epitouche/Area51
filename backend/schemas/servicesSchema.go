package schemas

import "time"

type Service struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
	Description		string		`json:"description" gorm:"type:varchar(100)"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt		time.Time	`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
