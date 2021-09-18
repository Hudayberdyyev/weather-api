package repository

import (
	"github.com/Hudayberdyyev/weather_api/models"
	"github.com/jackc/pgx"
)

type Forecast interface {
	GetCities() (*[]models.Regions, error)
	Create(regionId int, forecast *models.OwmResponse) error
}

type Repository struct {
	Forecast
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Forecast: NewForecastPostgres(db),
	}
}
