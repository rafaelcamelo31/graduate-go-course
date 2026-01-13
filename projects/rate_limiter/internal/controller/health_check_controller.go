package controller

import (
	"context"
	"net/http"

	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/usecase"
)

type HealthCheckController struct {
	healthCheckUseCase usecase.HealthCheckUseCaseInterface
}

func NewHealthCheckController(healthCheckUseCase usecase.HealthCheckUseCaseInterface) *HealthCheckController {
	return &HealthCheckController{
		healthCheckUseCase: healthCheckUseCase,
	}
}

func (hc *HealthCheckController) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	hc.healthCheckUseCase.GetHealthCheck(context.Background(), w, r)
}
