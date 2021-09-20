package main

import (
	"time"

	"github.com/Hudayberdyyev/weather_api/configs"
	"github.com/Hudayberdyyev/weather_api/pkg/repository"
	"github.com/Hudayberdyyev/weather_api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TickerHours = 2
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := configs.Init(); err != nil {
		logrus.Fatalf("error initializing configs: %s\n", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     uint16(viper.GetInt("db.port")),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error with initialize db: %s\n", err.Error())
	}

	defer db.Close()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	ticker := time.NewTicker(time.Duration(TickerHours) * time.Hour)

	for ; true; <-ticker.C {
		if err := services.Forecast.Update(); err != nil {
			logrus.Fatalf("error with update forecasts: %s", err.Error())
		}
	}

}
