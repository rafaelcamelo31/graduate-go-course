package mock

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type RateLimiterRepositoryMock struct {
	mock.Mock
}

func (m *RateLimiterRepositoryMock) IsBlocked(ctx context.Context, key string) (bool, error) {
	args := m.Called(ctx, key)
	return args.Bool(0), args.Error(1)
}

func (m *RateLimiterRepositoryMock) GetCount(ctx context.Context, key string) (int64, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RateLimiterRepositoryMock) Increment(ctx context.Context, key string) (int64, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RateLimiterRepositoryMock) SetCount(ctx context.Context, key string, window time.Duration) error {
	args := m.Called(ctx, key, window)
	return args.Error(0)
}

func (m *RateLimiterRepositoryMock) Block(ctx context.Context, key string, duration time.Duration) error {
	args := m.Called(ctx, key, duration)
	return args.Error(0)
}
