package main

import (
	"fmt"
	"sync"
	"time"
)

// 💡 Another example with channels

func worker(id int, tasks<- chan int,results chan<- int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("🟠 Worker %d starting\n",id)
	time.Sleep(2*time.Second)
	for task:=range tasks{
	results<-task*2
	}
	
	fmt.Printf("🟤 Worker %d finished!\n",id)
}

func main() {
	// create worker group
	var wg sync.WaitGroup
	numOfWorkers:= 3
	numOfJobs:= 5
	results:=make(chan int, numOfJobs)
	tasks:= make(chan int, numOfJobs)

	wg.Add(numOfWorkers)

	for i:= range numOfWorkers{
		go worker(i+1,tasks,results,&wg)
	}


	for i:= range numOfJobs{
		tasks <- i+1
	}

	close(tasks)

	go func() {
		wg.Wait() // Non blocking - We want to receive the vals in realtime
		close(results)

	}()

	// print the results
	for result:=range results{
		fmt.Println("✅ Result:",result)
	}

	fmt.Println("⭐ All workers finished ⭐")


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