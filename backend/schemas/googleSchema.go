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

type GoogleActionResponse struct {
	Id                 uint64   `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	User               User     `json:"user,omitempty" gorm:"foreignkey:UserId;references:Id"`
	UserId             uint64   `json:"-"`
	Worflow            Workflow `json:"workflow,omitempty" gorm:"foreignkey:WorkflowId;references:Id"`
	WorkflowId         uint64   `json:"-"`
	ResultSizeEstimate int      `json:"result_size_estimate"`
}

type GoogleActionOptionsInfo struct {
	ResultSizeEstimate int `json:"resultSizeEstimate"`
}
