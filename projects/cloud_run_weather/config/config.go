package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type WeatherConfig struct {
	ApiKey string
}

func LoadConfig() (*WeatherConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("error loading config", err)
		return nil, err
	}

	wc := &WeatherConfig{
		ApiKey: os.Getenv("WEATHER_API_KEY"),
	}

	return wc, nil
}
