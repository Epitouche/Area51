package schemas

type GoogleAction string

const (
	GoogleGetEmailAction GoogleAction = "get_email_action"
)

type GoogleReaction string

const ()

type GoogleResponseToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type GoogleUserInfo struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type GoogleActionOptions struct {
	Label string `json:"label"`
}
