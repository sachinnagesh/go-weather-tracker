package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/sachinnagesh/go-weather-tracker/service"
)

func GetCityWeather(w http.ResponseWriter, r *http.Request) {

	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	log.Println(city)
	data, err := service.Query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)

}
