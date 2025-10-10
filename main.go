package main

import (
	"fmt"
	"sync"
	"time"
)

// "Fixed Window Counter" implementation ⚡
// Why mutex? - So that we can protect our data when we're modifying it (locking/unlocking the critical section).

type RateLimiter struct {
	mu sync.Mutex
	count int
	limit int
	window time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter{
	return &RateLimiter{
		limit: limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow()bool{
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now:=time.Now()

	if now.After(rl.resetTime){
		rl.resetTime= now.Add(rl.window)
		rl.count = 0
	}

	if rl.count<rl.limit{
		rl.count++
		return true
	}
	return false
}

func main() {
	rateLimiter:= NewRateLimiter(5,2*time.Second) // 5 requests

	for range 10{
		if rateLimiter.Allow(){
			fmt.Println("Request Allowed ✅")
		}else{
			fmt.Println("Request Denied ❌")
		}
		time.Sleep(200 * time.Millisecond) // some delay
	}

}

// OUTPUT:
// $ go run .
// Request Allowed ✅
// Request Allowed ✅
// Request Allowed ✅
// Request Allowed ✅
// Request Allowed ✅
// Request Denied ❌
// Request Denied ❌
// Request Denied ❌
// Request Denied ❌
// Request Denied ❌
