package core_entity

import (
	"context"
	"time"
)

type RateLimiter struct {
	MaxRequest    int
	Window        time.Duration
	BlockDuration time.Duration
}

func NewRateLimiter(maxRequest int, window, blockDuration time.Duration) *RateLimiter {
	return &RateLimiter{
		MaxRequest:    maxRequest,
		Window:        window,
		BlockDuration: blockDuration,
	}
}

type RateLimiterInterface interface {
	IsBlocked(ctx context.Context, key string) (bool, error)
	GetCount(ctx context.Context, key string) (int64, error)
	Increment(ctx context.Context, key string) (int64, error)
	SetCount(ctx context.Context, key string, window time.Duration) error
	Block(ctx context.Context, key string, blockDuration time.Duration) error
}
