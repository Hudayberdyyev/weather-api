package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

var owmURL string

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
		owmURL = URL
		owmURL = strings.Replace(owmURL, "{city name}", value.Name, 1)
		owmURL = strings.Replace(owmURL, "{API key}", APIKey, 1)
		urlParsed, err := url.Parse(owmURL)
		fmt.Println(urlParsed.String())

		spaceClient := http.Client{
			Timeout: time.Second * 5, // Timeout after 2 seconds
		}
		req, err := http.NewRequest(http.MethodGet, urlParsed.String(), nil)
		if err != nil {
			return err
		}

		req.Header.Set("Accept", "*/*")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")

		res, getErr := spaceClient.Do(req)
		if getErr != nil {
			return getErr
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return readErr
		}

		var result models.OwmResponse

		jsonErr := json.Unmarshal(body, &result)
		if jsonErr != nil {
			return jsonErr
		}

		var firstTimestamp int64

		if len(result.ListData) > 0 {
			firstTimestamp = result.ListData[0].Dt
		} else {
			return errors.New("empty response from OpenWeather")
		}

		if err = s.repo.DeleteOld(value.RegionId, firstTimestamp); err != nil {
			return err
		}

		if err = s.repo.Create(value.RegionId, &result); err != nil {
			return err
		}
		logrus.Println(value, "ok")
	}

	if err = s.repo.DeleteEarlyCurrentDate(); err != nil {
		return err
	}
	return nil
}
