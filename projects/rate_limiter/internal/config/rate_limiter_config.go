package config

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	core_entity "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/entity"
)

const (
	API_KEYS         = "API_KEYS"
	API_KEY          = "API_KEY"
	IP               = "IP"
	MAX_REQUEST      = "_MAX_REQUEST"
	WINDOW           = "_WINDOW"
	BLOCK_DURATION   = "_BLOCK_DURATION"
	TOO_MANY_REQUEST = "you have reached the maximum number of requests or actions allowed within a certain time frame"
)

type RateLimiterConfig struct {
	Name          string
	ExtractKey    func(r *http.Request) (string, bool)
	LimiterPerKey func(key string) (*core_entity.RateLimiter, bool)
}

func BuildRateLimiterConfig() ([]RateLimiterConfig, error) {
	apiKeyLimiters, err := loadApiKeyLimiters()
	if err != nil {
		return nil, err
	}

	ipLimiter, err := loadRateLimiter(IP)
	if err != nil {
		return nil, err
	}

	return []RateLimiterConfig{
		{
			Name: API_KEY,
			ExtractKey: func(r *http.Request) (string, bool) {
				key := r.Header.Get(API_KEY)
				if key == "" {
					return "", false
				}
				return key, true
			},
			LimiterPerKey: func(key string) (*core_entity.RateLimiter, bool) {
				limiter, ok := apiKeyLimiters[key]
				return limiter, ok
			},
		},
		{
			Name: IP,
			ExtractKey: func(r *http.Request) (string, bool) {
				ip, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					return "", false
				}
				return ip, true
			},
			LimiterPerKey: func(key string) (*core_entity.RateLimiter, bool) {
				return ipLimiter, true
			},
		},
	}, nil
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

func loadApiKeyLimiters() (map[string]*core_entity.RateLimiter, error) {
	keys := strings.Split(os.Getenv(API_KEYS), ",")
	result := make(map[string]*core_entity.RateLimiter)

	for _, key := range keys {
		limiter, err := loadRateLimiter(API_KEY + "_" + key)
		if err != nil {
			return nil, err
		}
		result[key] = limiter
	}

	return result, nil
}
