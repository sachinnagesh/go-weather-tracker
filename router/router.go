package router

import (
	"net/http"

	"github.com/sachinnagesh/go-weather-tracker/handler"
)

func CreateRouter() {
	http.HandleFunc("/hello", handler.Hello)
	http.HandleFunc("/weather/", handler.GetCityWeather)
}
