package schemas

type InterpolAction string

const (
)

type InterpolReaction string

const (
	InterpolGetRedNotices      InterpolAction = "get_red_notices"
	InterpolGetYellowNotices   InterpolAction = "get_yellow_notices"
	InterpolGetUNNotices       InterpolAction = "get_un_notices"
)

type Warrants struct {
	Charge            string `json:"charge"`
	IssuingCountryId  string `json:"issuing_country_id"`
	ChargeTranslation string `json:"charge_translation"`
}

type InterpolNotice struct {
	ArrestWarrants      *[]Warrants `json:"arrest_warrants"`
	Weight              *uint32     `json:"weight"`
	Forename            *string     `json:"forename"`
	DateOfBirth         *string     `json:"date_of_birth"`
	EntityId            *string     `json:"entity_id"`
	LanguagesSpokenIds  *[]string   `json:"languages_spoken_ids"`
	Nationalities       *[]string   `json:"nationalities"`
	Height              *uint32     `json:"height"`
	SexId               *string     `json:"sex_id"`
	CountryOfBirthId    *string     `json:"country_of_birth_id"`
	Name                *string     `json:"name"`
	DistinguishingMarks *string     `json:"distinguishing_marks"`
	EyesColorsId        *string     `json:"eyes_colors_id"`
	HairsId             *string     `json:"hairs_id"`
	PlaceOfBirth        *string     `json:"place_of_birth"`
}

type InterpolEmbedded struct {
	Notices []InterpolNotice `json:"notices"`
}

type InterpolNoticesList struct {
	Total uint64 `json:"total"`
	Embedded InterpolEmbedded `json:"_embedded"`
}

type InterpolReactionOption struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type InterpolReactionOptionInfos struct {
	IsOld    bool           `json:"is_old"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

type InterpolRedNoticesInfos struct {
	Total uint64 `json:"total"`
}
