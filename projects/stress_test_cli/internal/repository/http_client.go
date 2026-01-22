package repository

import (
	"context"
	"net/http"
	"time"
)

type HTTPClientRepository interface {
	SendRequest(ctx context.Context, url string) (statusCode int, duration int64, err error)
}

type httpClientRepository struct {
	client *http.Client
}

func NewHTTPClientRepository(timeout time.Duration) HTTPClientRepository {
	client := &http.Client{
		Timeout: timeout,
	}
	return &httpClientRepository{
		client: client,
	}
}

func (h *httpClientRepository) SendRequest(ctx context.Context, url string) (int, int64, error) {
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, 0, err
	}

	resp, err := h.client.Do(req)
	duration := time.Since(start).Milliseconds()

	if err != nil {
		return 0, duration, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, duration, nil
}
