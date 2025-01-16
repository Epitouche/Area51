package schemas

type InterpolAction string

const (
	InterpolNewNotices InterpolAction = "new_notices"
)

type InterpolReaction string

const (

)

type InterpolActionOption struct {
	Total uint64 `json:"total"`
	IsOld bool   `json:"is_old"`
}

type InterpolRedNoticesInfos struct {
	Total uint64 `json:"total"`
}