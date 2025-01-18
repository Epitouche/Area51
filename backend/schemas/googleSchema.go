package schemas

type GoogleAction string

const (
	GoogleGetEmailAction GoogleAction = "get_email_action"
)

type GoogleReaction string

const (
	GoogleCreateEventReaction GoogleReaction = "create_event_reaction"
)

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

type GoogleCalendarCorpusOptionsTime struct {
	DateTime string `json:"dateTime"`
	TimeZone string `json:"timeZone"`
}

type GoogleCalendarCorpusOptionsTimeStartSchema struct {
	StartDateTime string `json:"startDateTime"`
	StartTimeZone string `json:"startTimeZone"`
}

type GoogleCalendarCorpusOptionsTimeEndSchema struct {
	EndDateTime string `json:"endDateTime"`
	EndTimeZone string `json:"endTimeZone"`
}

type GoogleCalendarCorpusOptionsSchema struct {
	Summary     string                                     `json:"summary"`
	Description string                                     `json:"description"`
	Location    string                                     `json:"location"`
	Start       GoogleCalendarCorpusOptionsTimeStartSchema `json:"start"`
	End         GoogleCalendarCorpusOptionsTimeEndSchema   `json:"end"`
	Attendees   GoogleCalendarCorpusOptionsAttendees       `json:"attendees"`
}

type GoogleCalendarOptionsSchema struct {
	CalendarId     string                            `json:"calendar_id"`
	CalendarCorpus GoogleCalendarCorpusOptionsSchema `json:"calendar_corpus"`
}

type GoogleCalendarCorpusOptionsAttendees struct {
	Email string `json:"email"`
}

type GoogleCalendarCorpusOptions struct {
	Summary     string                               `json:"summary"`
	Description string                               `json:"description"`
	Location    string                               `json:"location"`
	Start       GoogleCalendarCorpusOptionsTime      `json:"start"`
	End         GoogleCalendarCorpusOptionsTime      `json:"end"`
	Attendees   GoogleCalendarCorpusOptionsAttendees `json:"attendees"`
}

type GoogleCalendarOptions struct {
	CalendarId     string                      `json:"calendar_id"`
	CalendarCorpus GoogleCalendarCorpusOptions `json:"calendar_corpus"`
}

type GoogleCalendarResponse struct {
	Items []struct {
		Id string `json:"id"`
	} `json:"items"`
}
