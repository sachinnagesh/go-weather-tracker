package main

import (
	"net/http"

	"github.com/sachinnagesh/go-weather-tracker/router"
)

func main() {
	router.CreateRouter()

	http.ListenAndServe(":8080", nil)
}
