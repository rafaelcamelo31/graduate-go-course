package adapter

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/temperature_service/internal/entity"
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
		slog.Error("error in weather api request", "error", err, "request", req)
		return nil, err
	}

	resp, err := ha.client.Do(req)
	if err != nil {
		slog.Error("error in weather api response", "error", err, "response", resp)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading body", "error", err, "body", string(body))
		return nil, err
	}
	if strings.Contains(string(body), "error") {
		slog.Error("no matching location found")
		return nil, nil
	}

	wapi := *entity.NewWeather()
	err = json.Unmarshal(body, &wapi)
	if err != nil {
		return nil, err
	}
	slog.Info("successfully unmarshaled weather api", "data", wapi)

	return &wapi, nil
}
