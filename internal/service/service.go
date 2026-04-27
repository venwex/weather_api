package service

import "weather_api/internal/repository"
import "weather_api/internal/client"

type Service struct {
	User    *UserService
	City    *CityService
	Weather *WeatherService
}

func NewService(repos *repository.Repository, weatherClient client.WeatherClient) *Service {
	return &Service{
		User:    NewUserService(repos.Users),
		City:    NewCityService(repos.Users, repos.Cities),
		Weather: NewWeatherService(repos.Users, repos.Cities, repos.WeatherHistory, weatherClient),
	}
}
