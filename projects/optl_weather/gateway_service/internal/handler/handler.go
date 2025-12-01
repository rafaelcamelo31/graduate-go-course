package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/gateway_service/internal/adapter"
	"github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/gateway_service/internal/entity"
	apierror "github.com/rafaelcamelo31/graduate-go-course/projects/optl_weather/gateway_service/internal/error"
)

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
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

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	cep := city.Cep
	temp, err := adapter.GetTemperatureAdapter(ctx, cep)
	if err != nil {
		apierror.InternalServerError(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)
}
