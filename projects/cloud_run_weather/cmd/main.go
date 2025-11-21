package main

import (
	"log"
	"net/http"

	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/internal/handler"
)

func main() {
	http.HandleFunc("GET /api/weather", handler.GetWeatherHandler)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
