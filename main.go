package main

import (
	"fmt"
	"sync"
	"time"
)

// wait groups (sync pkg) -  wait for a collection of goroutines to complete their execution (another mechanism apart from channels)
// why? - synchronization, coordination, resource management
// basic ops - Add(delta int), Done(), Wait()

func worker(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("🔵 Worker %d starting\n",id)
	time.Sleep(time.Second) // simulate some time spent of processing this task
	fmt.Printf("✅ Worker %d finished!\n",id)
}

func main() {
	// create worker group
	var wg sync.WaitGroup
	numOfWorkers:= 3

	wg.Add(numOfWorkers)

	// Launch workers
	for i:= range numOfWorkers{
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("☑️ All workers finished!")


	// Output:
	// $ go run .
	// 🔵 Worker 2 starting
	// 🔵 Worker 1 starting
	// 🔵 Worker 0 starting
	// ✅ Worker 0 finished!
	// ✅ Worker 2 finished!
	// ✅ Worker 1 finished!
	// ☑️ All workers finished!
}