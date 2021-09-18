package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Hudayberdyyev/weather_api/models"
	"github.com/Hudayberdyyev/weather_api/pkg/repository"
)

const (
	APIKey = "559cd760f10f8731db7e748be0666c37"
	URL    = "http://api.openweathermap.org/data/2.5/forecast?q={city name}&appid={API key}&lang=ru&units=metric"
)

var url string

type ForecastService struct {
	repo repository.Forecast
}

func NewForecastService(repo repository.Forecast) *ForecastService {
	return &ForecastService{repo: repo}
}

func (s *ForecastService) Update() error {
	cities, err := s.repo.GetCities()

	if err != nil {
		return err
	}

	for _, value := range *cities {
		url = URL
		url = strings.Replace(url, "{city name}", value.Name, 1)
		url = strings.Replace(url, "{API key}", APIKey, 1)
		fmt.Println(url)

		spaceClient := http.Client{
			Timeout: time.Second * 5, // Timeout after 2 seconds
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			logrus.Fatal(err)
		}

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			logrus.Fatal(getErr)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			logrus.Fatal(readErr)
		}

		var result models.OwmResponse

		jsonErr := json.Unmarshal(body, &result)
		if jsonErr != nil {
			logrus.Fatal(jsonErr)
		}

		if err = s.repo.Create(value.RegionId, &result); err != nil {
			return err
		}
		logrus.Println(value, "ok")
	}

	return nil
}
