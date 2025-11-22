package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/config"
	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/internal/adapter"
	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/internal/handler"
)

func main() {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	cfg := &config.WeatherConfig{}
	cfg.LoadConfig()

	viacepAdapter := adapter.NewHttpViaCEPAdapter(httpClient)
	weatherAdapter := adapter.NewHttpWeatherApiAdapter(httpClient, cfg.ApiKey)

	h := handler.NewHandler(viacepAdapter, weatherAdapter, cfg)
	http.HandleFunc("GET /api/temperature", h.GetTemperature)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
