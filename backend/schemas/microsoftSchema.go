package schemas

type MicrosoftAction string

const (
	MicrosoftOutlookEventsAction MicrosoftAction = "get_outlook_events"
)

type MicrosoftReaction string

const ()

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

type MicrosoftOutlookEventsOptions struct {
	Subject string `json:"subject"`
}

type MicrosoftOutlookEventsResponse struct {
	Value []struct {
		Subject string `json:"subject"`
	} `json:"value"`
}
