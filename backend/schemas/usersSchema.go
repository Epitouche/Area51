package schemas

type User struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
	LastName		string		`json:"lastname" gorm:"type:varchar(100)"`
	Username		string		`json:"username" gorm:"type:varchar(100);unique"`
	Email			string		`json:"email" binding:"requiredcredentials" gorm:"type:varchar(100);unique"`
	Password		string		`json:"password" gorm:"type:varchar(100)"` // can be null for Oauth2 user
	Image			string		`json:"image"`
	IsAdmin			bool		`json:"is_admin" gorm:"type:boolean"`
	UserRole		uint64		`json:"user_role"`

	UserRoleID		Role		`json:"user_role_id" gorm:"foreignkey:UserRole;references:Id"`
}
