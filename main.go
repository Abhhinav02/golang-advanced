package main

import (
	"fmt"
	"time"
)

// "Token Bucket Algorithm" implementation ⚡
// Why struct{}{} ? - No memory-overhead with an empty struct{} (0 bytes) 💡

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

 func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter{
	rl:= &RateLimiter{
	tokens: make(chan struct{}, rateLimit),
	refillTime: refillTime,
}
	for range rateLimit{
	rl.tokens <- struct{}{}
}
	go rl.startRefill()
	return  rl
 }

 func (rl *RateLimiter) startRefill(){
	ticker:= time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <-struct{}{}:
			default:
			}
		}
	}
} 

func (rl *RateLimiter) allow() bool{
	select {
	case <-rl.tokens:
		return true
	default:
		return false	
	}
}

func main() {
	rateLimiter:= NewRateLimiter(5,time.Second) // 5 requests

	// Let's send 10 requests
	for range 10{
		if rateLimiter.allow(){
			fmt.Println("Request Allowed ✅")
		}else{
			fmt.Println("Request denied ❌")
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
// Request Allowed ✅
// Request denied ❌
// Request denied ❌
// Request denied ❌
// Request denied ❌