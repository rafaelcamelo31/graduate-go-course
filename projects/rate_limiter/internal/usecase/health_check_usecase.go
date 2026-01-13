package usecase

import (
	"context"
	"encoding/json"
	"net/http"
)

type HealthCheckUseCase struct{}

type HealthCheckUseCaseInterface interface {
	GetHealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

func NewHealthCheckUseCase() *HealthCheckUseCase {
	return &HealthCheckUseCase{}
}

type Response struct {
	Message string
	Status  int
}

func (hc *HealthCheckUseCase) GetHealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{
		Message: "Server responded with success.",
		Status:  http.StatusOK,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
