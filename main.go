package main

import (
	"fmt"
	"time"
)

// timer in action - sending the current time after a certain period
func main() {
	fmt.Println("Starting app..")
	timer := time.NewTimer(2 * time.Second) // Non-blocking in nature (unlike time.Sleep())
	fmt.Println("Waiting for timer.C")
	stopped:=timer.Stop()
	if stopped{
		fmt.Println("Timer stopped..")
	}
	fmt.Println("Timer reset()..")
	timer.Reset(time.Second)
	<- timer.C // blocking in nature
	fmt.Println("Timer expired!")

	//OP:
    // $ go run .
    // Starting app..
    // Waiting for timer.C
    // Timer stopped..
    // Timer reset()..
    // Timer expired!

}