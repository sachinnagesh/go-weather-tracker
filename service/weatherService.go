package service

import (
	"encoding/json"
	"net/http"

	"github.com/sachinnagesh/go-weather-tracker/entity"
)

func Query(city string) (entity.WeatherData, error) {
	apiconfig, err := entity.LoadAPIConfig("./resources/.apiConfig")
	if err != nil {
		return entity.WeatherData{}, err
	}
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiconfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return entity.WeatherData{}, err
	}

	defer res.Body.Close()
	var d entity.WeatherData
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return entity.WeatherData{}, err
	}
	return d, nil
}
