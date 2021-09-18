package main

import (
	"github.com/Hudayberdyyev/weather_api/configs"
	"github.com/Hudayberdyyev/weather_api/pkg/repository"
	"github.com/Hudayberdyyev/weather_api/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TickerHours = 12
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

	// ticker := time.NewTicker(time.Duration(TickerHours) * time.Hour)

	// for _ = range ticker.C {
	if err := services.Forecast.Update(); err != nil {
		logrus.Fatalf("error with update forecasts: %s", err.Error())
	}
	// }

	// for _, value := range cities {
	// url := URL
	// url = strings.Replace(url, "{city name}", value, 1)
	// url = strings.Replace(url, "{API key}", APIKey, 1)
	// fmt.Println(url)

	// spaceClient := http.Client{
	// 	Timeout: time.Second * 5, // Timeout after 2 seconds
	// }

	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, getErr := spaceClient.Do(req)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }

	// if res.Body != nil {
	// 	defer res.Body.Close()
	// }

	// body, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }

	// var result models.OwmResponse

	// jsonErr := json.Unmarshal(body, &result)
	// if jsonErr != nil {
	// 	log.Fatal(jsonErr)
	// }

	// 	fmt.Println(value, "ok")
	// }
}
