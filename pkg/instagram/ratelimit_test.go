package instagram

import (
	"net/http"
	"testing"
	"time"
)

func TestRateLimiterPacing(t *testing.T) {
	limiter := NewRateLimiter(200 * time.Millisecond)
	start := time.Now()

	if err := limiter.Throttle(); err != nil {
		t.Fatalf("expected no error on first throttle call, got: %v", err)
	}

	if err := limiter.Throttle(); err != nil {
		t.Fatalf("expected no error on second throttle call, got: %v", err)
	}

	elapsed := time.Since(start)
	if elapsed < 200*time.Millisecond {
		t.Errorf("expected throttle to enforce at least 200ms delay, got %v", elapsed)
	}
}

func TestRateLimiterCircuitBreaker(t *testing.T) {
	limiter := NewRateLimiter(50 * time.Millisecond)
	resp := &http.Response{StatusCode: http.StatusTooManyRequests}

	limiter.HandleResponse(resp, nil)
	limiter.HandleResponse(resp, nil)
	limiter.HandleResponse(resp, nil)

	summary := limiter.GetStatusSummary()
	if summary["circuit_breaker_active"] != true {
		t.Errorf("expected circuit breaker to be active after 3 consecutive 429s")
	}
}
