package main

import (
	"fmt"
	"sync"
)

/*
💡mutexes - mitual exclution, sync. primitive which prevents multiple goroutines from simultaneously accessing shared resources or exeuting critical sections of the code. Ensures that only 1 goroutine can hold the 'mutex' at a time, thus avaoiding race-conditions and data-corruption. 💻
*/

type Counter struct{
	mu sync.Mutex
	val int
}

func (c *Counter)increment(){
	c.mu.Lock()
	c.val++
	defer c.mu.Unlock()
}

func (c *Counter) getVal()int{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func main() {
	// waitgroups for multiple goroutines
	var wg sync.WaitGroup
	counter:= &Counter{}
	numOfGoroutines := 10

	for range numOfGoroutines{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100{
				counter.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("✅ Final counter val: %d\n",counter.getVal()) // ⚠️unreliable output without 'mutex'

	// O.P- 
	// $ go run .
	// ✅ Final counter val: 1000
}