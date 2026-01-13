package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/controller"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/middleware"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/usecase"
)

const (
	APP_HOST = "APP_HOST"
	APP_PORT = "APP_PORT"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	useCase := usecase.NewHealthCheckUseCase()
	c := controller.NewHealthCheckController(useCase)

	mux := http.NewServeMux()
	mux.Handle("GET /api/health", http.HandlerFunc(c.GetHealthCheck))

	log.Println("Starting server at:", os.Getenv(APP_PORT))
	http.ListenAndServe(os.Getenv(APP_HOST)+":"+os.Getenv(APP_PORT), middleware.RateLimiterMiddleware(mux))
}
