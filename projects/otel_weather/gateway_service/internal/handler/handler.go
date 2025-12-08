package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/gateway_service/internal/adapter"
	"github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/gateway_service/internal/entity"
	apierror "github.com/rafaelcamelo31/graduate-go-course/projects/otel_weather/gateway_service/internal/error"
)

type Handler struct {
	temperature adapter.TemperatureServiceInterface
}

func NewHandler(temperature adapter.TemperatureServiceInterface) *Handler {
	return &Handler{
		temperature: temperature,
	}
}

func (h *Handler) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	city := &entity.City{}
	err := json.NewDecoder(r.Body).Decode(city)
	if err != nil {
		apierror.BadRequestError(w)
		return
	}

	if !city.IsAllDigits() || !city.IsEightDigits() {
		apierror.InvalidCepError(w)
		return
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	cep := city.Cep
	temp, err := h.temperature.GetTemperatureAdapter(ctx, cep)
	if err != nil {
		apierror.InternalServerError(w)
		return
	}
	if temp == nil {
		apierror.NotFoundError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)
}
