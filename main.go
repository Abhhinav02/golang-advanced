package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Atomic counter - atomic ops.
// Cruicial for safely handling shared data in concurrent programming.
// More efficeint than mutexes. But use them for simple ops/counting/tracking... For complex ones, use 'mutexes'.

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getVal() int64{
	return atomic.LoadInt64(&ac.count)
}



func main() {
	var wg sync.WaitGroup
	numOfGoroutines:=10

	counter:= &AtomicCounter{}
	// value:=0 // ❌ unreliable

	for range numOfGoroutines{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000{
				counter.increment()
				// value++ // ❌ unreliable
			}
		}()
	}

	wg.Wait()
	fmt.Printf("✅ Final value: %d\n",counter.getVal())
	// fmt.Printf("✅ Final value: %d\n",value) // ❌ unreliable result

	//O.P - ✅ Final value: 10000
}