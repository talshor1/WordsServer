package rate_limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu            sync.Mutex
	allowedPerSec int
	allowedPerMin int
	lastSecTime   time.Time
	lastMinTime   time.Time
	requestsSec   int
	requestsMin   int
}

func NewRateLimiter(allowedPerSec, allowedPerMin int) *RateLimiter {
	return &RateLimiter{
		allowedPerSec: allowedPerSec,
		allowedPerMin: allowedPerMin,
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now()

	if now.Sub(r.lastSecTime) >= time.Second {
		r.lastSecTime = now
		r.requestsSec = 0
	}

	if now.Sub(r.lastMinTime) >= time.Minute {
		r.lastMinTime = now
		r.requestsMin = 0
	}

	if r.requestsSec < r.allowedPerSec && r.requestsMin < r.allowedPerMin {
		r.requestsSec++
		r.requestsMin++
		return true
	}

	return false
}
