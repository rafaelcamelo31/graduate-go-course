package config

import (
	"errors"
	"net/url"
)

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func NewConfig(url string, requests, concurrency int) (*Config, error) {
	config := &Config{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}

	if err := config.IsValid(); err != nil {
		return nil, err
	}

	return config, nil
}

type Result struct {
	StatusCode   int
	ResponseTime int64
	Error        error
}

type StressTestReport struct {
	TargetURL          string
	TotalRequests      int
	ConcurrencyLevel   int
	TotalElapsedTime   float64
	SuccessfulRequests int
	FailedRequests     int
}

func (c *Config) IsValid() error {
	if err := c.validateURL(c.URL); err != nil {
		return errors.New("invalid URL")
	}

	if c.Requests < 0 {
		return errors.New("requests must be greater than 0")
	}

	if c.Concurrency <= 0 {
		return errors.New("concurrency must be greater than or equal to 0")
	}

	if c.Concurrency > c.Requests {
		return errors.New("concurrency cannot be greater than requests")
	}

	return nil
}

func (c *Config) validateURL(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	return err
}
