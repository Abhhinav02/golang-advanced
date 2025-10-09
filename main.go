package main

import (
	"fmt"
	"sync"
	"time"
)

// Real-World scenario simulation 🚛

type Worker struct{
	ID int
	Task string
}

// PerformTask - worker simulation
func (w *Worker) PerformTask(wg *sync.WaitGroup){
 defer wg.Done()
 fmt.Printf("🚧 Worker %d started %s ...\n",w.ID,w.Task)
 time.Sleep(2*time.Second) // time taken to complete task
 fmt.Printf("✅ Worker %d finished %s\n",w.ID,w.Task)
}

func main() {
	// create waitgroup
	var wg sync.WaitGroup

	// define tasks to be performed by workers
	tasks:= []string{"Digging ⛏️","Laying bricks 🧱", "Painting 🖌️"} 

	for i,task := range tasks{
		worker := Worker{ID:i+1, Task: task}
		wg.Add(1) // can add in a loop too
		go worker.PerformTask(&wg)
	}

	// wait for all workers to finish
	wg.Wait()

	// Construction is finished
	fmt.Println("Construction completed.. ☑️")


	// 💻Output:
	// $ go run .
	// 🚧 Worker 3 started Painting 🖌️ ...
	// 🚧 Worker 1 started Digging ⛏️ ...
	// 🚧 Worker 2 started Laying bricks 🧱 ...
	// ✅ Worker 1 finished Digging ⛏️
	// ✅ Worker 2 finished Laying bricks 🧱
	// ✅ Worker 3 finished Painting 🖌️
	// Construction completed.. ☑️
	
}