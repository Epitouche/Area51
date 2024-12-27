package schemas

type UserServiceLink struct {
	Id        uint64  `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId    User    `json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	ServiceId Service `json:"service,omitempty" gorm:"foreignkey:ServiceId;references:Id"`
}
