package main

// LEAKY BUCKET ALGO. + GOROUTINES 💦

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens int
	lastLeak time.Time
	mu sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket{
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens: capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow()bool{
	lb.mu.Lock()
	defer lb.mu.Unlock()
	now:= time.Now()
	elapsedTime:= now.Sub(lb.lastLeak)
	tokensToAdd:= int(elapsedTime/lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity{
		lb.tokens=lb.capacity
	}
	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)


	fmt.Printf("🔥 Tokens added %d, Tokens subtracted %d, Total Tokens %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("🕰️ Last leak-time: %v", lb.lastLeak)

	if lb.tokens > 0{
		lb.tokens--
		return true
	}
	return false
}

func main() {
	leakyBucket:= NewLeakyBucket(5,500*time.Millisecond)
	var wg sync.WaitGroup

	for range 10{
		wg.Add(1)
		go func() {
			defer wg.Done()
			if leakyBucket.Allow(){
			fmt.Println("🕛 Curr. time:",time.Now())
			fmt.Println("Request Accepted ✅")
		}else{
			fmt.Println("🕛 Curr. time:",time.Now())
			fmt.Println("Request Rejected ❌")
		}
		time.Sleep(200*time.Millisecond)
		}()
		
	}

	time.Sleep(500*time.Millisecond)
	wg.Wait()
}

// O:P -
// go run .
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 5
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=+0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6825644 +0530 IST m=+0.002029101
// Request Accepted ✅
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 4
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=+0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6836974 +0530 IST m=+0.003162101
// Request Accepted ✅
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 3
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=+0.000745601🕛 Curr. time: 2025-10-11 12:25:42.684388 +0530 IST m=+0.003852701
// Request Accepted ✅
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 2
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=+0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6857004 +0530 IST m=+0.005165101
// Request Accepted ✅
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 1
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=+0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6857004 +0530 IST m=+0.005165101
// Request Accepted ✅
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 0
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=+0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6862702 +0530 IST m=+0.005734901
// Request Rejected ❌
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 0      
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=++0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6867999 +0530 IST m=+0.006264601
// Request Rejected ❌
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 0      
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=++0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6867999 +0530 IST m=+0.006264601
// Request Rejected ❌
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 0      
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=++0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6867999 +0530 IST m=+0.006264601
// Request Rejected ❌
// 🔥 Tokens added 0, Tokens subtracted 1, Total Tokens 0      
// 🕰️ Last leak-time: 2025-10-11 12:25:42.6812809 +0530 IST m=++0.000745601🕛 Curr. time: 2025-10-11 12:25:42.6873807 +0530 IST m=+0.006845401
// Request Rejected ❌

