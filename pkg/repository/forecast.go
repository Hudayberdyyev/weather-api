package repository

import (
	"fmt"
	"time"

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

func (r *ForecastPostgres) Create(regionId int, forecast *models.OwmResponse) error {
	var unixTimeUTC time.Time
	for _, value := range forecast.ListData {
		unixTimeUTC = time.Unix(value.Dt, 0)
		query := fmt.Sprintf("insert into %s (\"temp\", feels_like, owm_id, humidity, wind_speed, pop, pressure, ts, region_id) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", forecastsTable)
		if _, err := r.db.Exec(query, value.MainData.Temp, value.MainData.FeelsLike, value.WeatherData[0].ID, value.MainData.Humidity, value.WindData.Speed, value.Pop, value.MainData.Pressure, unixTimeUTC, regionId); err != nil {
			return err
		}
	}
	return nil
}

func (r *ForecastPostgres) DeleteOld(regionId int, ts int64) error {
	unixTimeUTC := time.Unix(ts, 0)
	query := fmt.Sprintf("delete from %s where ts > $1 and region_id = $2", forecastsTable)
	if _, err := r.db.Exec(query, unixTimeUTC, regionId); err != nil {
		return err
	}
	return nil
}

func (r *ForecastPostgres) DeleteEarlyCurrentDate() error {
	currentTime := time.Now()
	query := fmt.Sprintf("delete from %s where ts < $1", forecastsTable)
	if _, err := r.db.Exec(query, currentTime); err != nil {
		return err
	}
	return nil
}
