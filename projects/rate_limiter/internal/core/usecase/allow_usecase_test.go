package core_usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	core_entity "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/entity"
	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllow_KeyBlocked(t *testing.T) {
	repo := new(mock.RateLimiterRepositoryMock)
	usecase := NewRateLimiterUseCase(repo)

	ctx := context.Background()
	key := "API_KEY:blocked"
	limiter := &core_entity.RateLimiter{}

	repo.
		On("IsBlocked", ctx, key).
		Return(true, nil).
		Once()

	allowed, err := usecase.Allow(ctx, key, limiter)

	require.NoError(t, err)
	assert.False(t, allowed)

	repo.AssertExpectations(t)
}

func TestAllow_FirstRequestWithinLimit(t *testing.T) {
	repo := new(mock.RateLimiterRepositoryMock)
	usecase := NewRateLimiterUseCase(repo)

	ctx := context.Background()
	key := "IP:127.0.0.1"

	limiter := &core_entity.RateLimiter{
		MaxRequest:    5,
		Window:        time.Minute,
		BlockDuration: time.Minute * 10,
	}

	repo.On("IsBlocked", ctx, key).Return(false, nil).Once()
	repo.On("Increment", ctx, key).Return(int64(1), nil).Once()
	repo.On("SetCount", ctx, key, limiter.Window).Return(nil).Once()

	allowed, err := usecase.Allow(ctx, key, limiter)

	require.NoError(t, err)
	assert.True(t, allowed)

	repo.AssertExpectations(t)
}

func TestAllow_ExceedsLimit_ShouldBlock(t *testing.T) {
	repo := new(mock.RateLimiterRepositoryMock)
	usecase := NewRateLimiterUseCase(repo)

	ctx := context.Background()
	key := "API_KEY:limit"
	limiter := &core_entity.RateLimiter{
		MaxRequest:    2,
		Window:        time.Minute,
		BlockDuration: time.Minute * 5,
	}

	repo.On("IsBlocked", ctx, key).Return(false, nil).Once()
	repo.On("Increment", ctx, key).Return(int64(3), nil).Once()
	repo.On("Block", ctx, key, limiter.BlockDuration).Return(nil).Once()

	allowed, err := usecase.Allow(ctx, key, limiter)

	require.NoError(t, err)
	assert.False(t, allowed)

	repo.AssertExpectations(t)
}

func TestAllow_RepositoryError(t *testing.T) {
	repo := new(mock.RateLimiterRepositoryMock)
	usecase := NewRateLimiterUseCase(repo)

	ctx := context.Background()
	key := "error-key"
	limiter := &core_entity.RateLimiter{}

	repo.
		On("IsBlocked", ctx, key).
		Return(false, errors.New("db error")).
		Once()

	allowed, err := usecase.Allow(ctx, key, limiter)

	assert.Error(t, err)
	assert.False(t, allowed)

	repo.AssertExpectations(t)
}
