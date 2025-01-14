package schemas

type MicrosoftUserInfo struct {
	Mail        string `json:"mail"`
	DisplayName string `json:"displayName"`
}

type MicrosoftResponseToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}
