package adapter

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/gateway_service/internal/entity"
)

func GetTemperatureAdapter(ctx context.Context, cep string) (*entity.Temperature, error) {
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("error in sending temperature service request", "error", err)
		return nil, err
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
