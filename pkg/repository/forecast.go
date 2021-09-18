package repository

import (
	"fmt"

	"github.com/Hudayberdyyev/weather_api/models"
	"github.com/jackc/pgx"
)

type ForecastPostgres struct {
	db *pgx.Conn
}

func NewForecastPostgres(db *pgx.Conn) *ForecastPostgres {
	return &ForecastPostgres{db: db}
}

func (r *ForecastPostgres) GetCities() (*[]models.Regions, error) {
	var res []models.Regions
	query := fmt.Sprintf("select regions_id, title from %s where hl='en'", regionsTextTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var region models.Regions
		if err = rows.Scan(&region.RegionId, &region.Name); err != nil {
			return nil, err
		}
		res = append(res, region)
	}
	return &res, nil
}
