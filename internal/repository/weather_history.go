package repository

import (
	"context"
	"errors"
	"weather_api/internal/models"

	"github.com/jmoiron/sqlx"
)

type WeatherHistoryRepo struct {
	db *sqlx.DB
}

func NewWeatherHistoryRepository(db *sqlx.DB) *WeatherHistoryRepo {
	return &WeatherHistoryRepo{
		db: db,
	}
}

func (r *WeatherHistoryRepo) CreateWeatherHistory(ctx context.Context, item models.WeatherHistory) (models.WeatherHistory, error) {
	return models.WeatherHistory{}, errors.New("not implemented")
}

func (r *WeatherHistoryRepo) GetWeatherHistoryByCity(ctx context.Context, userID int, city string, limit int) ([]models.WeatherHistory, error) {
	return nil, errors.New("not implemented")
}
