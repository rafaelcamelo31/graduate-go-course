package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/config"
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
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}
	addr := os.Getenv(APP_HOST) + ":" + os.Getenv(APP_PORT)

	redisClient := config.NewRedisClient()

	repo := repository.NewRedisRepository(redisClient)
	usecase := core_usecase.NewRateLimiterUseCase(repo)
	c := controller.NewHealthCheckController()

	mux := http.NewServeMux()
	mux.Handle("GET /api/health", http.HandlerFunc(c.GetHealthCheck))

	log.Println("Starting server at:", os.Getenv(APP_PORT))

	http.ListenAndServe(addr, middleware.RateLimiterMiddleware(mux, usecase))
}
