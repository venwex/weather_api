package handler

import "weather_api/internal/service"

type Handler struct {
	Users   *UserHandler
	Cities  *CityHandler
	Weather *WeatherHandler
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		Users:   NewUserHandler(svc.User),
		Cities:  NewCityHandler(svc.City),
		Weather: NewWeatherHandler(svc.Weather),
	}
}
