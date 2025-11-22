package adapter

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/rafaelcaemlo31/graduate-go-course/projects/cloud_run_weather/internal/entity"
)

type ViaCepAdapterInterface interface {
	GetCityByCep(ctx context.Context, cep string) (*entity.City, error)
}

var _ ViaCepAdapterInterface = (*HttpViaCEPAdapter)(nil)

type HttpViaCEPAdapter struct {
	client *http.Client
	URL    string
}

func NewHttpViaCepAdapter(client *http.Client) *HttpViaCEPAdapter {
	return &HttpViaCEPAdapter{
		client: client,
		URL:    "https://viacep.com.br/ws/",
	}
}

func (ha *HttpViaCEPAdapter) GetCityByCep(ctx context.Context, cep string) (*entity.City, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", ha.URL+cep+"/json/", nil)
	if err != nil {
		slog.Error("error in viacep request", "error", err, "request", req)
		return nil, err
	}

	resp, err := ha.client.Do(req)
	if err != nil {
		slog.Error("error in viacep response", "error", err, "response", resp)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading data from body", "error", err)
		return nil, err
	}
	if strings.Contains(string(body), "erro") {
		return nil, nil
	}

	city := *entity.NewCity(cep)
	err = json.Unmarshal(body, &city)
	if err != nil {
		slog.Error("error unmarshalling viacep to city", "error", err, "city", city)
		return nil, err
	}

	return &city, nil
}
