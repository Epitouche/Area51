package schemas

type Token struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId			User		`json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	ServiceId		Service		`json:"service,omitempty" gorm:"foreignkey:ServiceId;references:Id"`
	Token			string		`json:"token" gorm:"type:varchar(100)"`
}
