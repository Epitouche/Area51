package api

import "area51/controllers"

type WeatherApi struct {
	controller controllers.WeatherController
}

func NewWeatherApi(controller controllers.WeatherController) *WeatherApi {
	return &WeatherApi{
		controller: controller,
	}
}
