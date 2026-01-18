package middleware

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/config"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/controller"
	core_usecase "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/usecase"
)

const (
	TOO_MANY_REQUEST = "you have reached the maximum number of requests or actions allowed within a certain time frame"
)

func RateLimiterMiddleware(next http.Handler, usecase core_usecase.RateLimiterInterface, config []config.RateLimiterConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		for _, cfg := range config {
			key, ok := cfg.ExtractKey(r)
			if !ok {
				continue
			}

			limiter, ok := cfg.LimiterPerKey(key)
			if !ok {
				http.Error(w, "invalid API key", http.StatusBadRequest)
				return
			}

			allowed, err := usecase.Allow(ctx, cfg.Name+":"+key, limiter)
			if err != nil {
				slog.Error("rate limiter error", "rate_limiter", cfg.Name, "error", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}

			if !allowed {
				sendTooManyRequest(w)
				return
			}

			break
		}

		next.ServeHTTP(w, r)
	})
}

func sendTooManyRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTooManyRequests)
	response := &controller.Response{
		Message: TOO_MANY_REQUEST,
		Status:  http.StatusTooManyRequests,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	slog.Info(TOO_MANY_REQUEST)
}
