package main

import (
	"fmt"
	"sync"
)

//💡mutexes - how do they understand which values to protect?

func main() {

	var counter int
	var wg sync.WaitGroup

	var mu sync.Mutex

	numOfGoroutines:=5
	wg.Add(numOfGoroutines)

	increment:= func ()  {
		// can be used inside loops too
		defer wg.Done()
		for range 1000{
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numOfGoroutines{
		go increment()
	}

	wg.Wait()
	fmt.Printf("✅ Final counter val: %d\n",counter)
	
	// O.P- 
	// $ go run .
	// ✅ Final counter val: 5000
	
}