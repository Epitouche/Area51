package controllers

type WeatherController interface {
}

type weatherController struct {
}

func NewWeatherController() WeatherController {
	return &weatherController{}
}
