package service

import "weather_api/internal/repository"

type Service struct {
	User *UserService
	City *CityService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.Users),
		City: NewCityService(repos.Users, repos.Cities),
	}
}
