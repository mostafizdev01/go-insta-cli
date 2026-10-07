package instagram

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"go-insta-cli/pkg/cli"
)

// RateLimiter manages account protection controls and request throttling.
type RateLimiter struct {
	mu                    sync.Mutex
	lastRequestTime       time.Time
	minInterval           time.Duration
	consecutiveRateLimits int
	isCircuitOpen         bool
	cooldownUntil         time.Time
}

// GlobalLimiter is the package instance for Meta Graph API rate control.
var GlobalLimiter = NewRateLimiter(1500 * time.Millisecond)

// NewRateLimiter creates a new rate limiter with a minimum safety interval.
func NewRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		minInterval: interval,
	}
}

// Throttle enforces human-like request pacing and circuit breaker checks.
func (rl *RateLimiter) Throttle() error {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	// Check Circuit Breaker status
	if rl.isCircuitOpen {
		if now.Before(rl.cooldownUntil) {
			remaining := time.Until(rl.cooldownUntil).Round(time.Second)
			return fmt.Errorf("Rate limit detected. Circuit breaker active. Pausing requests for %v", remaining)
		}
		// Reset circuit breaker after cooldown
		rl.isCircuitOpen = false
		rl.consecutiveRateLimits = 0
	}

	// Calculate safety delay interval
	elapsed := now.Sub(rl.lastRequestTime)
	if elapsed < rl.minInterval {
		sleepTime := rl.minInterval - elapsed
		time.Sleep(sleepTime)
	}

	rl.lastRequestTime = time.Now()
	return nil
}

// HandleResponse inspects HTTP response status for rate limits or HTTP 429.
func (rl *RateLimiter) HandleResponse(resp *http.Response, err error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		rl.consecutiveRateLimits++
		fmt.Printf("\n%s⚠ Rate limit detected (HTTP 429). Applying exponential backoff.%s\n", cli.ColorYellow, cli.ColorReset)

		// Exponential Backoff calculation
		backoffDuration := time.Duration(1<<rl.consecutiveRateLimits) * 5 * time.Second
		if backoffDuration > 60*time.Second {
			backoffDuration = 60 * time.Second
		}

		if rl.consecutiveRateLimits >= 3 {
			rl.isCircuitOpen = true
			rl.cooldownUntil = time.Now().Add(60 * time.Second)
			fmt.Printf("%s⚠ Circuit Breaker Tripped! Pausing API requests for account safety for 60s.%s\n\n", cli.ColorRed, cli.ColorReset)
		} else {
			rl.cooldownUntil = time.Now().Add(backoffDuration)
		}
		return
	}

	if resp != nil && resp.StatusCode == http.StatusOK {
		if rl.consecutiveRateLimits > 0 {
			rl.consecutiveRateLimits = 0
			rl.isCircuitOpen = false
		}
	}
}

// GetStatusSummary returns formatted rate limit status information.
func (rl *RateLimiter) GetStatusSummary() map[string]interface{} {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	statusStr := "Normal (Account Safe)"
	if rl.isCircuitOpen {
		statusStr = "Circuit Open (Pausing)"
	} else if rl.consecutiveRateLimits > 0 {
		statusStr = "Throttled (Backoff Active)"
	}

	return map[string]interface{}{
		"status":                 statusStr,
		"safety_interval_ms":     rl.minInterval.Milliseconds(),
		"consecutive_rate_limits": rl.consecutiveRateLimits,
		"circuit_breaker_active": rl.isCircuitOpen,
	}
}
