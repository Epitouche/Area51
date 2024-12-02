package schemas

type PasswordRecovery struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	UserId			uint64		`json:"user_id,omitempty"`
	Code			string		`json:"code" gorm:"type:varchar(100)"`
	IsValidated		bool		`json:"is_validated" gorm:"type:boolean"`

	User			User		`json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
}
