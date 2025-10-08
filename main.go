package main

import (
	"fmt"
	"time"
)

// worker-pool : Design pattern to manage a group of workers (goroutines)
// used for resource-management

func worker(id int, tasks<-chan int, results chan<- int){
	for task := range tasks{
		fmt.Printf("Worker %d processing task %d\n",id,task)
		// Simulate work/ops.
		time.Sleep(time.Second)
		results <-task*2
	}
}


func main() {
	numOfWorkers := 3
	numOfJobs:= 10
	tasks:= make(chan int,numOfJobs)
	results:= make(chan int,numOfJobs)

	// Create workers
	for i:=range numOfWorkers{
		go worker(i,tasks,results)
	}

	// Send values to the tasks channel
	for i:=range numOfJobs{
		tasks<-i
	}

	close(tasks)

	// Collect the results
	for range numOfJobs{
		result:=<-results
		fmt.Println("Result:",result)
	}
}