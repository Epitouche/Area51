package schemas

type WeatherAction string

const (
	WeatherCurrentAction WeatherAction = "current_weather"
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

type WeatherCurrentOptions struct {
	CityName     string  `json:"city_name"`
	LanguageCode string  `json:"language_code"`
	Temperature  float64 `json:"temperature"`
	CompareSign  string  `json:"compare_sign"`
}
