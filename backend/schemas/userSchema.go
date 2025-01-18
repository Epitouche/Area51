package schemas

type User struct {
	Id       uint64         `json:"id,omitempty" gorm:"primary_key;auto_increment;"`
	Name     string         `json:"name" gorm:"type:varchar(100)"`
	LastName string         `json:"lastname" gorm:"type:varchar(100)"`
	Username string         `json:"username" gorm:"type:varchar(100);unique"`
	Email    *string        `json:"email" gorm:"type:varchar(100);unique"`
	Password *string        `json:"password" gorm:"type:varchar(100)"`
	Image    string         `json:"image" gorm:"type:BYTEA"`
	IsAdmin  bool           `json:"is_admin" gorm:"type:boolean"`
	Services []ServiceToken `gorm:"many2many:user_service_tokens;constraint:OnDelete:CASCADE;"`
}
