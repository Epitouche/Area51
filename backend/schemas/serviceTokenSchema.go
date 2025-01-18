package schemas

import "time"

type ServiceToken struct {
	Id           uint64    `gorm:"primaryKey;autoIncrement"           json:"id,omitempty"`
	UserId       uint64    `                                          json:"-"`
	User         User      `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"    json:"user_id"`
	ServiceId    uint64    `                                          json:"-"`
	Service      Service   `gorm:"foreignKey:ServiceId;references:Id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"service_id"`
	Token        string    `                                          json:"token"`
	RefreshToken string    `                                          json:"refresh_token"`
	ExpireAt     time.Time `                                          json:"expireAt"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"createdAt"`
	UpdateAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"          json:"updateAt"`
	Users        []User    `gorm:"many2many:user_service_tokens;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
