package service

import "github.com/Hudayberdyyev/weather_api/pkg/repository"

type Forecast interface {
	Update() error
}

type Service struct {
	Forecast
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Forecast: NewForecastService(repo.Forecast),
	}
}
