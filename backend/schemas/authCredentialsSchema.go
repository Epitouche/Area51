package schemas

type LoginCredentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterCredentials struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

type MobileToken struct {
	Token string `json:"token"`
}

type GithubCodeCredentials struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
