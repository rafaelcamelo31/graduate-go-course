package main

import (
	"log"
	"net/http"
	"os"
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

	viacepAdapter := adapter.NewHttpViaCepAdapter(httpClient)
	weatherAdapter := adapter.NewHttpWeatherApiAdapter(httpClient, cfg.ApiKey)

	h := handler.NewHandler(viacepAdapter, weatherAdapter, cfg)
	http.HandleFunc("GET /api/temperature", h.GetTemperature)

	port := os.Getenv("PORT")
	log.Printf("Starting server on %s", port)
	http.ListenAndServe(":"+port, nil)
}
