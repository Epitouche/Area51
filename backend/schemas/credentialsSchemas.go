package schemas

type RegisterCredentials struct {
	Email string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type JWT struct {
	Token string `json:"token"`
}