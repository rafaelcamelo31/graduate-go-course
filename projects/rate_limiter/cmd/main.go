package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/controller"
	core_usecase "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/usecase"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/middleware"
	repository "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/repository/redis"
)

const (
	APP_HOST = "APP_HOST"
	APP_PORT = "APP_PORT"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	port := os.Getenv(APP_PORT)
	if port == "" {
		log.Fatal("APP_PORT is not set")
	}

	redisClient := config.NewRedisClient()

	repo := repository.NewRedisRepository(redisClient)
	usecase := core_usecase.NewRateLimiterUseCase(repo)
	c := controller.NewHealthCheckController()

	mux := http.NewServeMux()
	mux.Handle("GET /api/health", http.HandlerFunc(c.GetHealthCheck))

	config, err := config.BuildRateLimiterConfig()
	if err != nil {
		log.Fatal("rate limiter config caused an error", err)
	}

	slog.Info("Starting server at:", "port", port)
	if err := http.ListenAndServe(":"+port, middleware.RateLimiterMiddleware(mux, usecase, config)); err != nil {
		log.Fatal(err)
	}
}
