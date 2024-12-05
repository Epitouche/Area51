package schemas

type Role struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name			string		`json:"name" gorm:"type:varchar(100)"`
}
