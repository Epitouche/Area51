package schemas

import "time"

type OAuth2Token struct {
	Id				uint64		`json:"id,omitempty" gorm:"primary_key;auto_increment"`
	AccessToken		string		`json:"access_token"`
	Scope			string		`json:"scope"`
	TokenType		string		`json:"token_type"`
	CreatedAt		time.Time	`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type GithubTokenReponse struct {
	AccesToken 		string 		`json:"access_token"`
	Scope 			string 		`json:"scope"`
	TokenType 		string 		`json:"token_type"`
}
