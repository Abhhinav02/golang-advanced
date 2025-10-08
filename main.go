package main

import (
	"fmt"
	"time"
)

// time.After() - After waits for the duration to elapse and then sends the current time on the returned channel.
func main() {
	timeOut:=time.After(3*time.Second)
	done:= make(chan bool)

	go func ()  {
		longRunningOp()
		done <-true
	}()

	select{
	case <-timeOut:
		fmt.Println("Operation timed out!.. 🔴")
	case <-done:
		fmt.Println("✅ Operation completed/done!")	
	}

	// OP:
	// $ go run .
	// 0
	// 1
	// 2
	// Operation timed out!.. 🔴
	

}

// simulating a resource-heavy/time-consuming func()
func longRunningOp(){
	for i:=range 20{
       fmt.Println(i)
	   time.Sleep(time.Second) // sleep after every rep.
	}
}