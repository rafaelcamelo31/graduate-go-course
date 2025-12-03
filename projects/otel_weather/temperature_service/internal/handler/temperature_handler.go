package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/constant"
	apierror "github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/error"
	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/internal/adapter"
	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/temperature_service/internal/entity"
)

type Handler struct {
	viacepAdapter  adapter.ViaCepAdapterInterface
	weatherAdapter adapter.WeatherApiAdapterInterface
	config         *config.WeatherConfig
}

func NewHandler(viacepAdapter adapter.ViaCepAdapterInterface, weatherAdapter adapter.WeatherApiAdapterInterface, cfg *config.WeatherConfig) *Handler {
	return &Handler{
		viacepAdapter:  viacepAdapter,
		weatherAdapter: weatherAdapter,
		config:         cfg,
	}
}

func (h *Handler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cep := r.URL.Query().Get("cep")

	if cep == "" {
		http.Error(w, constant.MISSING_CEP_QUERY, http.StatusBadRequest)
		slog.Error(constant.MISSING_CEP_QUERY, "code", http.StatusBadRequest)
		return
	}

	c := *entity.NewCity(cep)
	if !c.IsEightDigits() || !c.IsAllDigits() {
		http.Error(w, constant.INVALID_ZIPCODE, http.StatusUnprocessableEntity)
		slog.Error(constant.INVALID_ZIPCODE, "code", http.StatusUnprocessableEntity)
		return
	}

	city, err := h.viacepAdapter.GetCityByCep(ctx, cep)
	if err != nil {
		apierror.GetInternalServerError(w, err)
		return
	}
	if city == nil {
		http.Error(w, constant.CANNOT_FIND_CEP, http.StatusNotFound)
		slog.Error(constant.CANNOT_FIND_CEP, "cep", cep, "status", http.StatusNotFound)
		return
	}

	weather, err := h.weatherAdapter.GetWeather(ctx, city.Name)
	if err != nil {
		apierror.GetInternalServerError(w, err)
		return
	}
	if weather == nil {
		http.Error(w, constant.CANNOT_FIND_WEATHER, http.StatusNotFound)
		slog.Error(constant.CANNOT_FIND_WEATHER, "weather", weather, "status", http.StatusNotFound)
		return
	}

	t := entity.NewTemperature(weather.Current.TempC, weather.Current.Tempf, city.Name)
	slog.Info("Temperature in", "city", city, "temperature", fmt.Sprintf("%.1fC", weather.Current.TempC))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}
