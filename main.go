package main

import (
	"fmt"
	"time"
)

// Scheduling delayed operations.
func main() {
	timer:=time.NewTimer(2*time.Second) // non-blocking timer
	go func() {
		<- timer.C
		fmt.Println("✅ Delayed Op. executed ⌛")
	}()

	fmt.Println("Waiting.. ")
	time.Sleep(3*time.Second) // blocking timer
	fmt.Println("End of the program.. ☑️")

	// OP:
	// 	$ go run .
	// Waiting.. 
	// ✅ Delayed Op. executed ⌛
	// End of the program.. ☑️

}
