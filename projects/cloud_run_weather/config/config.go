package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigInterface interface {
	LoadConfig() (*WeatherConfig, error)
}

var _ ConfigInterface = (*WeatherConfig)(nil)

type WeatherConfig struct {
	ApiKey string
}

func (wc *WeatherConfig) LoadConfig() (*WeatherConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("error loading config", err)
		return nil, err
	}

	wc.ApiKey = os.Getenv("WEATHER_API_KEY")

	return wc, nil
}
