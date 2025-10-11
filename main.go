package main

import (
	"fmt"
	"time"
)

// stateful groutines - state is preserved between contexts
// COUNTER example

type StatefulWorker struct{
	count int
	ch chan int
}

// Receivig value from ch.
func(sw *StatefulWorker) Start(){
	go func() {
		// infinite loop
		for {
			select {
			case value:= <-sw.ch:
				sw.count+=value
				fmt.Println("✅ Curr. count:",sw.count)
				
			}
		}
	}()
}

// Sending value to ch.
func(sw *StatefulWorker) Send(value int){
	sw.ch <- value
}

func main() {
	stWorker:= &StatefulWorker{
		ch: make(chan int),
	}
	stWorker.Start()

	for i:=range 5{
		stWorker.Send(i)
		time.Sleep(500*time.Millisecond)
	}
}

// O.P - 
// ✅ Curr. count: 0
// ✅ Curr. count: 1
// ✅ Curr. count: 3
// ✅ Curr. count: 6
// ✅ Curr. count: 10