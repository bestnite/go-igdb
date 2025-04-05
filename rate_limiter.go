package igdb

import (
	"sync"
	"time"
)

type rateLimiter struct {
	mu         sync.Mutex
	rate       int
	interval   time.Duration
	tokens     int
	lastRefill time.Time
}

func newRateLimiter(rate int) *rateLimiter {
	return &rateLimiter{
		rate:       rate,
		interval:   time.Second,
		tokens:     rate,
		lastRefill: time.Now(),
	}
}

func (r *rateLimiter) wait() {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefill)

	if elapsed >= r.interval {
		r.tokens = r.rate
		r.lastRefill = now
	}

	if r.tokens <= 0 {
		waitTime := r.interval - elapsed
		r.mu.Unlock()
		time.Sleep(waitTime)
		r.mu.Lock()
		r.tokens = r.rate - 1
		r.lastRefill = time.Now()
		return
	}

	r.tokens--
}
