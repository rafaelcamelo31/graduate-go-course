package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/internal/adapter"
	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/internal/handler"
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
