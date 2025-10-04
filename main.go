package main

import (
	"fmt"
	"time"
)

// NON-BLOCKING OPS.

func main() {

	// non-blocking Ops. in REAL-TIME-SYSTEMS
	dataCh := make(chan int)
	quitCh := make(chan bool)

	go func(){
		for {
			select {
			case d:=<-dataCh:
				fmt.Println("✅ Data received:",d)
			
			case <-quitCh:
				fmt.Println("Stopping... 🔴")
				return
			default:
				fmt.Println("Waiting for data... ⏳")
				time.Sleep(500 * time.Millisecond)
		}
	}
	}()

	for i:=range 5{
		dataCh<-i
		time.Sleep(time.Second)
	}

	quitCh<-true


}