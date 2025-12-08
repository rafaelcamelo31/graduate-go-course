package adapter

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/gateway_service/internal/entity"
)

type TemperatureServiceInterface interface {
	GetTemperatureAdapter(ctx context.Context, cep string) (*entity.Temperature, error)
}

var _ TemperatureServiceInterface = (*HttpTemperatureServiceAdapter)(nil)

type HttpTemperatureServiceAdapter struct {
	client *http.Client
}

func NewHttpTemperatureServiceAdapter(client *http.Client) *HttpTemperatureServiceAdapter {
	return &HttpTemperatureServiceAdapter{
		client: client,
	}
}

func (ht *HttpTemperatureServiceAdapter) GetTemperatureAdapter(ctx context.Context, cep string) (*entity.Temperature, error) {
	u := &url.URL{
		Scheme: "http",
		Host:   "temperature-api:8081",
		Path:   "/api/temperature",
	}
	q := u.Query()
	q.Add("cep", cep)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		slog.Error("error in creating temperature service request", "error", err)
		return nil, err
	}

	resp, err := ht.client.Do(req)
	if err != nil {
		slog.Error("error in sending temperature service request", "error", err)
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		slog.Error(http.StatusText(http.StatusNotFound), "status", http.StatusNotFound)
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error in reading response body", "error", err)
		return nil, err
	}

	temp := &entity.Temperature{}
	err = json.Unmarshal(body, temp)
	if err != nil {
		slog.Error("error in unmarshalling body into temperature", "error", err)
		return nil, err
	}

	return temp, nil
}
