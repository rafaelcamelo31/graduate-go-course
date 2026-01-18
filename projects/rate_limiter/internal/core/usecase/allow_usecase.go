package core_usecase

import (
	"context"
	"log/slog"

	core_entity "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/entity"
)

type RateLimiterInterface interface {
	Allow(ctx context.Context, key string, rateLimiter *core_entity.RateLimiter) (bool, error)
}

type RateLimiterUseCase struct {
	repository core_entity.RateLimiterInterface
}

func NewRateLimiterUseCase(repo core_entity.RateLimiterInterface) *RateLimiterUseCase {
	return &RateLimiterUseCase{
		repository: repo,
	}
}

func (ak *RateLimiterUseCase) Allow(ctx context.Context, key string, rateLimiter *core_entity.RateLimiter) (bool, error) {
	slog.Info("checking if the key is blocked")

	blocked, err := ak.repository.IsBlocked(ctx, key)
	if err != nil {
		slog.Error("error at repository", "error", err)
		return false, err
	}
	if blocked {
		slog.Info("the key is blocked", "key", key)
		return false, nil
	}

	slog.Info("incrementing (or creating) key value...")
	count, err := ak.repository.Increment(ctx, key)
	if err != nil {
		slog.Error("error at repository", "error", err)
		return false, err
	}

	slog.Info("setting max request allowed in rateLimiter window...", "window", rateLimiter.Window)
	if count == 1 {
		if err := ak.repository.SetCount(ctx, key, rateLimiter.Window); err != nil {
			slog.Error("error at repository", "error", err)
			return false, err
		}
	}

	slog.Info("checking if the count exceeded max request")
	if count > int64(rateLimiter.MaxRequest) {
		if err := ak.repository.Block(ctx, key, rateLimiter.BlockDuration); err != nil {
			slog.Error("error at repository", "error", err)
			return false, err
		}
		return false, nil
	}

	slog.Info("request successfully passed rate limiter")
	return true, nil
}
