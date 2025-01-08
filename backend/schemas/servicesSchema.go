package schemas

import "time"

type ServiceName string

const (
	Github ServiceName = "github"
)

type ServiceJson struct {
	Name        ServiceName    `json:"name"`
	Description string         `json:"description"`
	Action      []ActionJson   `json:"actions"`
	Reaction    []ReactionJson `json:"reactions"`
	Image       string         `json:"image"`
}

type Service struct {
	Id          uint64      `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name        ServiceName `json:"name" gorm:"type:varchar(100)"`
	Description string      `json:"description" gorm:"type:varchar(100)"`
	Image       string      `json:"image" gorm:"type:BYTEA"`
	CreatedAt   time.Time   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
