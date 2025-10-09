package main

import (
	"fmt"
	"sync"
	"time"
)

// 💡 WAITGROUPS WITH CHANNELS 🔥

func worker(id int, results chan<- int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("🟡 Worker %d starting\n",id)
	time.Sleep(2*time.Second) // simulate some time spent on processing this task
	results<-id*2 // results ch. receive some values
	fmt.Printf("🟣 Worker %d finished!\n",id)
}

func main() {
	// create worker group
	var wg sync.WaitGroup
	numOfWorkers:= 3
	numOfJobs:=3
	results:=make(chan int, numOfJobs) // 3 buffers

	wg.Add(numOfWorkers)

	// Launch workers
	for i:= range numOfWorkers{
		go worker(i,results ,&wg)
	}

	go func() {
		wg.Wait() // Non blocking - We want to receive the vals in realtime
		close(results) 	// close the channel once all goroutines have finished

	}()

	// print the results
	for result:=range results{
		fmt.Println("✅ Result:",result)
	}

	//fmt.Println("☑️ All workers finished!")


	// Output:
	// $ go run .
	// 🟡 Worker 0 starting
	// 🟡 Worker 2 starting
	// 🟡 Worker 1 starting
	// 🟣 Worker 1 finished!
	// 🟣 Worker 0 finished!
	// 🟣 Worker 2 finished!
	// ✅ Result: 2
	// ✅ Result: 0
	// ✅ Result: 4
	

}