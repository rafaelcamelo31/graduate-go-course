package middleware

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	core_entity "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/entity"
	core_usecase "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/usecase"
)

const (
	API_KEY        = "API_KEY"
	IP             = "IP"
	MAX_REQUEST    = "_MAX_REQUEST"
	WINDOW         = "_WINDOW"
	BLOCK_DURATION = "_BLOCK_DURATION"
)

type Middleware struct {
	usecase core_usecase.RateLimiterInterface
}

func NewMiddleware(usecase core_usecase.RateLimiterInterface) *Middleware {
	return &Middleware{
		usecase: usecase,
	}
}

func RateLimiterMiddleware(next http.Handler, usecase core_usecase.RateLimiterInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestApiKey := r.Header.Get(API_KEY)
		apiKeyLimiter, err := loadRateLimiter(API_KEY)
		if err != nil {
			slog.Error("invalid API_KEY rate limiter config", "error", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		clientIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			slog.Error("error getting ip from request", "error", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		ipLimiter, err := loadRateLimiter(IP)
		if err != nil {
			slog.Error("invalid IP rate limiter config", "error", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		envApiKey := os.Getenv(API_KEY)

		ctx := context.Background()
		if requestApiKey == envApiKey {
			usecase.Allow(ctx, requestApiKey, apiKeyLimiter)
		} else {
			usecase.Allow(ctx, clientIp, ipLimiter)
		}

		next.ServeHTTP(w, r)
	})
}

func loadRateLimiter(limiterType string) (*core_entity.RateLimiter, error) {
	maxRequest, err := strconv.Atoi(os.Getenv(limiterType + MAX_REQUEST))
	if err != nil {
		slog.Error("error converting string MAX_REQUEST", "error", err)
		return nil, err
	}

	window, err := time.ParseDuration(os.Getenv(limiterType + WINDOW))
	if err != nil {
		slog.Error("error converting string WINDOW", "error", err)
		return nil, err
	}

	blockDuration, err := time.ParseDuration(os.Getenv(limiterType + BLOCK_DURATION))
	if err != nil {
		slog.Error("error converting string BLOCK_DURATION", "error", err)
		return nil, err
	}

	return &core_entity.RateLimiter{
		MaxRequest:    maxRequest,
		Window:        window,
		BlockDuration: blockDuration,
	}, nil
}
