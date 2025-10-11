package main

// LEAKY BUCKET ALGO. 💦

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


	fmt.Printf("⚡ Tokens added %d, Tokens subtracted %d, Total Tokens %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("🕰️ Last leak-time: %v", lb.lastLeak)

	if lb.tokens > 0{
		lb.tokens--
		return true
	}
	return false
}

func main() {
	leakyBucket:= NewLeakyBucket(5,500*time.Millisecond)

	for range 10{
		if leakyBucket.Allow(){
			fmt.Println("🕛 Curr. time:",time.Now())
			fmt.Println("Request Accepted ✅")
		}else{
			fmt.Println("🕛 Curr. time:",time.Now())
			fmt.Println("Request Rejected ❌")
		}
		time.Sleep(200*time.Millisecond)
	}
}

// O:P
// go run .
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 5
// 🕰️ Last leak-time: 2025-10-11 12:12:21.7196414 +0530 IST m=+0.000524001🕛 Curr. time: 2025-10-11 12:12:21.7219396 +0530 IST m=+0.002822201
// Request Accepted ✅
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 4
// 🕰️ Last leak-time: 2025-10-11 12:12:21.7196414 +0530 IST m=+0.000524001🕛 Curr. time: 2025-10-11 12:12:21.9240145 +0530 IST m=+0.204382901
// Request Accepted ✅
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 3
// 🕰️ Last leak-time: 2025-10-11 12:12:21.7196414 +0530 IST m=+0.000524001🕛 Curr. time: 2025-10-11 12:12:22.1252766 +0530 IST m=+0.406159201
// Request Accepted ✅
// ⚡ Tokens added 1, Tokens subtracted 1, Total Tokens 3
// 🕰️ Last leak-time: 2025-10-11 12:12:22.2196414 +0530 IST m=+0.500524001🕛 Curr. time: 2025-10-11 12:12:22.326646 +0530 IST m=+0.607528601
// Request Accepted ✅
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 2
// 🕰️ Last leak-time: 2025-10-11 12:12:22.2196414 +0530 IST m=+0.500524001🕛 Curr. time: 2025-10-11 12:12:22.528829 +0530 IST m=+0.809711601
// Request Accepted ✅
// ⚡ Tokens added 1, Tokens subtracted 1, Total Tokens 2
// 🕰️ Last leak-time: 2025-10-11 12:12:22.7196414 +0530 IST m=+1.000524001🕛 Curr. time: 2025-10-11 12:12:22.730006 +0530 IST m=+1.010888601
// Request Accepted ✅
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 1
// 🕰️ Last leak-time: 2025-10-11 12:12:22.7196414 +0530 IST  m=+1.000524001🕛 Curr. time: 2025-10-11 12:12:22.9319658 +0530 IST m=+1.212848401
// Request Accepted ✅
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 0
// 🕰️ Last leak-time: 2025-10-11 12:12:22.7196414 +0530 IST  m=+1.000524001🕛 Curr. time: 2025-10-11 12:12:23.1344171 +0530 IST m=+1.415299701
// Request Rejected ❌
// ⚡ Tokens added 1, Tokens subtracted 1, Total Tokens 1
// 🕰️ Last leak-time: 2025-10-11 12:12:23.2196414 +0530 IST  m=+1.500524001🕛 Curr. time: 2025-10-11 12:12:23.3355704 +0530 IST m=+1.616453001
// Request Accepted ✅
// ⚡ Tokens added 0, Tokens subtracted 1, Total Tokens 0
// 🕰️ Last leak-time: 2025-10-11 12:12:23.2196414 +0530 IST  m=+1.500524001🕛 Curr. time: 2025-10-11 12:12:23.5376747 +0530 IST m=+1.818557301
// Request Rejected ❌
