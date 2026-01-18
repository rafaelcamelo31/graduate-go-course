package stres_stest

import (
	"net/http"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestRateLimiter_APIKey_Stress(t *testing.T) {
	rate := vegeta.Rate{
		Freq: 100,
		Per:  time.Second,
	}
	duration := 10 * time.Second

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/api/health",
		Header: http.Header{
			"API_KEY": []string{"TEST_KEY_1"},
		},
	})

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, rate, duration, "rate-limiter") {
		metrics.Add(res)

		if res.Code != http.StatusOK && res.Code != http.StatusTooManyRequests {
			t.Fatalf("unexpected status: %d", res.Code)
		}
	}

	metrics.Close()

	t.Logf("Requests: %d\n", metrics.Requests)
	t.Logf("Success: %.2f%%\n", metrics.Success*100)
	t.Logf("429: %d\n", metrics.StatusCodes["429"])
}

func TestRateLimiter_Burst(t *testing.T) {
	rate := vegeta.Rate{
		Freq: 1000,
		Per:  time.Second,
	}
	duration := 3 * time.Second

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/api/health",
	})

	attacker := vegeta.NewAttacker(
		vegeta.Workers(200),
	)

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "burst") {
		metrics.Add(res)
	}

	metrics.Close()

	t.Logf("429 count: %d \n", metrics.StatusCodes["429"])
}

func TestRateLimiter_MultipleKeys(t *testing.T) {
	keys := []string{
		"TEST_KEY_1",
		"TEST_KEY_2",
		"TEST_KEY_3",
	}

	rate := vegeta.Rate{Freq: 300, Per: time.Second}
	duration := 10 * time.Second

	targeter := func() vegeta.Targeter {
		return func(t *vegeta.Target) error {
			key := keys[time.Now().UnixNano()%int64(len(keys))]
			t.Method = "GET"
			t.URL = "http://localhost:8080/api/health"
			t.Header = http.Header{
				"API_KEY": []string{key},
			}
			return nil
		}
	}()

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, rate, duration, "multiple-keys") {
		metrics.Add(res)
	}

	metrics.Close()

	t.Logf("status codes: %+v", metrics.StatusCodes)
}
