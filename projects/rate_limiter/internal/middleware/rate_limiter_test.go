package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/config"
	core_entity "github.com/rafaelcamelo31/graduate-go-course/projects/rate_limiter/internal/core/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type RateLimiterUseCaseMock struct {
	mock.Mock
}

func (m *RateLimiterUseCaseMock) Allow(
	ctx context.Context,
	key string,
	limiter *core_entity.RateLimiter,
) (bool, error) {
	args := m.Called(ctx, key, limiter)
	return args.Bool(0), args.Error(1)
}

func TestRateLimiterMiddleware_APIKeyAllowed(t *testing.T) {
	usecase := new(RateLimiterUseCaseMock)
	limiter := &core_entity.RateLimiter{}

	usecase.
		On("Allow", mock.Anything, "API_KEY:TEST_KEY_1", limiter).
		Return(true, nil).
		Once()

	cfg := []config.RateLimiterConfig{
		{
			Name: "API_KEY",
			ExtractKey: func(r *http.Request) (string, bool) {
				return r.Header.Get("API_KEY"), true
			},
			LimiterPerKey: func(key string) (*core_entity.RateLimiter, bool) {
				return limiter, true
			},
		},
	}

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	req.Header.Set("API_KEY", "TEST_KEY_1")
	rec := httptest.NewRecorder()

	handler := RateLimiterMiddleware(next, usecase, cfg)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, nextCalled)

	usecase.AssertExpectations(t)
}

func TestRateLimiterMiddleware_APIKeyBlocked(t *testing.T) {
	usecase := new(RateLimiterUseCaseMock)
	limiter := &core_entity.RateLimiter{}

	usecase.
		On("Allow", mock.Anything, "API_KEY:TEST_KEY_1", limiter).
		Return(false, nil).
		Once()

	cfg := []config.RateLimiterConfig{
		{
			Name: "API_KEY",
			ExtractKey: func(r *http.Request) (string, bool) {
				return "TEST_KEY_1", true
			},
			LimiterPerKey: func(key string) (*core_entity.RateLimiter, bool) {
				return limiter, true
			},
		},
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("next handler must not be called")
	})

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()

	handler := RateLimiterMiddleware(next, usecase, cfg)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusTooManyRequests, rec.Code)

	usecase.AssertExpectations(t)
}

func TestRateLimiterMiddleware_FallbackToIP(t *testing.T) {
	usecase := new(RateLimiterUseCaseMock)

	ipLimiter := &core_entity.RateLimiter{}

	usecase.
		On("Allow", mock.Anything, "IP:127.0.0.1", ipLimiter).
		Return(true, nil).
		Once()

	cfg := []config.RateLimiterConfig{
		{
			Name: "API_KEY",
			ExtractKey: func(r *http.Request) (string, bool) {
				return "", false
			},
			LimiterPerKey: func(key string) (*core_entity.RateLimiter, bool) {
				return nil, false
			},
		},
		{
			Name: "IP",
			ExtractKey: func(r *http.Request) (string, bool) {
				return "127.0.0.1", true
			},
			LimiterPerKey: func(key string) (*core_entity.RateLimiter, bool) {
				return ipLimiter, true
			},
		},
	}

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	rec := httptest.NewRecorder()

	handler := RateLimiterMiddleware(next, usecase, cfg)
	handler.ServeHTTP(rec, req)

	require.True(t, nextCalled)

	usecase.AssertExpectations(t)
}
