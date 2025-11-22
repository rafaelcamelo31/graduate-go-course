package adapter

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/internal/entity"
)

type WeatherApiAdapterInterface interface {
	GetWeather(ctx context.Context, city string) (*entity.Weather, error)
}

var _ WeatherApiAdapterInterface = (*HttpWeatherApiAdapter)(nil)

type HttpWeatherApiAdapter struct {
	client *http.Client
	apiKey string
	URL    string
}

func NewHttpWeatherApiAdapter(client *http.Client, apiKey string) *HttpWeatherApiAdapter {
	return &HttpWeatherApiAdapter{
		client: client,
		apiKey: apiKey,
		URL:    "https://api.weatherapi.com/v1/current.json",
	}
}

func (ha *HttpWeatherApiAdapter) GetWeather(ctx context.Context, city string) (*entity.Weather, error) {
	params := url.Values{}
	params.Add("key", ha.apiKey)
	params.Add("q", city)

	req, err := http.NewRequestWithContext(ctx, "GET", ha.URL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := ha.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		slog.Error("WeatherAPI returned status", "code", resp.StatusCode, "body", string(body))
		return nil, err
	}

	wapi := *entity.NewWeather()
	err = json.NewDecoder(resp.Body).Decode(&wapi)
	if err != nil {
		return nil, err
	}
	slog.Info("successfully unmarshaled weather api", "data", wapi)

	return &wapi, nil
}
