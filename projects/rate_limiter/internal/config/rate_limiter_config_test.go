package config

import (
	"net/http"
	"testing"
	"time"
)

func TestBuildRateLimiterConfig(t *testing.T) {
	t.Setenv(API_KEYS, "TEST_KEY_1,TEST_KEY_2")

	t.Setenv("API_KEY_TEST_KEY_1"+MAX_REQUEST, "5")
	t.Setenv("API_KEY_TEST_KEY_1"+WINDOW, "10s")
	t.Setenv("API_KEY_TEST_KEY_1"+BLOCK_DURATION, "30s")

	t.Setenv("API_KEY_TEST_KEY_2"+MAX_REQUEST, "10")
	t.Setenv("API_KEY_TEST_KEY_2"+WINDOW, "20s")
	t.Setenv("API_KEY_TEST_KEY_2"+BLOCK_DURATION, "40s")

	t.Setenv("IP"+MAX_REQUEST, "100")
	t.Setenv("IP"+WINDOW, "1m")
	t.Setenv("IP"+BLOCK_DURATION, "2m")

	configs, err := BuildRateLimiterConfig()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(configs) != 2 {
		t.Fatalf("expected 2 rate limiter configs, got %d", len(configs))
	}

	var apiKeyConfig, ipConfig *RateLimiterConfig
	for i := range configs {
		switch configs[i].Name {
		case API_KEY:
			apiKeyConfig = &configs[i]
		case IP:
			ipConfig = &configs[i]
		}
	}

	if apiKeyConfig == nil {
		t.Fatal("API_KEY config not found")
	}
	if ipConfig == nil {
		t.Fatal("IP config not found")
	}

	reqWithKey, _ := http.NewRequest("GET", "/api/health", nil)
	reqWithKey.Header.Set(API_KEY, "TEST_KEY_1")

	key, ok := apiKeyConfig.ExtractKey(reqWithKey)
	if !ok || key != "TEST_KEY_1" {
		t.Fatalf("expected API key 'TEST_KEY_1', got '%s'", key)
	}

	// TEST_KEY_1
	limiter, ok := apiKeyConfig.LimiterPerKey("TEST_KEY_1")
	if !ok {
		t.Fatal("expected limiter for TEST_KEY_1")
	}

	if limiter.MaxRequest != 5 {
		t.Errorf("expected MaxRequest=5, got %d", limiter.MaxRequest)
	}
	if limiter.Window != 10*time.Second {
		t.Errorf("expected Window=10s, got %v", limiter.Window)
	}
	if limiter.BlockDuration != 30*time.Second {
		t.Errorf("expected BlockDuration=30s, got %v", limiter.BlockDuration)
	}

	// TEST_KEY_2
	limiter2, ok := apiKeyConfig.LimiterPerKey("TEST_KEY_2")
	if !ok {
		t.Fatal("expected limiter for TEST_KEY_2")
	}
	if limiter2.MaxRequest != 10 {
		t.Errorf("expected MaxRequest=10, got %d", limiter2.MaxRequest)
	}

	_, ok = apiKeyConfig.LimiterPerKey("unknown")
	if ok {
		t.Fatal("expected no limiter for unknown api key")
	}

	// IP
	reqWithIP, _ := http.NewRequest("GET", "/api/health", nil)
	reqWithIP.RemoteAddr = "192.168.0.1:3131"

	ip, ok := ipConfig.ExtractKey(reqWithIP)
	if !ok || ip != "192.168.0.1" {
		t.Fatalf("expected IP '192.168.0.1', got '%s'", ip)
	}

	ipLimiter, ok := ipConfig.LimiterPerKey(ip)
	if !ok {
		t.Fatal("expected IP limiter")
	}
	if ipLimiter.MaxRequest != 100 {
		t.Errorf("expected IP MaxRequest=100, got %d", ipLimiter.MaxRequest)
	}
}
