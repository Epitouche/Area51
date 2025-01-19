package schemas

type WeatherAction string

const (
	WeatherCurrentAction WeatherAction = "current_feeling_temperature"
	WeatherTimeAction    WeatherAction = "sunrise_events"
)

type WeatherReaction string

const (
	WeatherCurrentReaction WeatherReaction = "current_weather"
)

type WeatherActionOptions struct {
	Current struct {
		Feelslike_c float64 `json:"feelslike_c"`
	} `json:"current"`
}

type WeatherReactionOptions struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Temp_c    float64 `json:"temp_c"`
		IsDay     int     `json:"is_day"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type WeatherCurrentOptions struct {
	CityName     string `json:"city_name"`
	LanguageCode string `json:"language_code"`
	Temperature  string `json:"temperature"`
	CompareSign  string `json:"compare_sign"`
}

type WeatherCurrentReactionOptions struct {
	CityName     string `json:"city_name"`
	LanguageCode string `json:"language_code"`
}

type WeatherSpecificTimeOption struct {
	CityName string `json:"city_name"`
	DateTime string `json:"dt"`
}

type WeatherSpecificTimeInfos struct {
	Astronomy struct {
		Astro struct {
			Sunrise string `json:"sunrise"`
		} `json:"astro"`
	} `json:"astronomy"`
}
