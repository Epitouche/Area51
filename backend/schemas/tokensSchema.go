package schemas

type Token struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId			uint64		`json:"user_id,omitempty"`
	ServiceId		uint64		`json:"service_id,omitempty"`
	Token			string		`json:"token" gorm:"type:varchar(100)"`

	User			User		`json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	Service			Service		`json:"service,omitempty" gorm:"foreignkey:ServiceId;references:Id"`
}
