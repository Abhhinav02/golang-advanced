package main

import "fmt"

// multiplexing - kinda like switch-case
func main() {
	ch1:= make(chan int)
	ch2:= make(chan int)

	//! Err - DEADLOCK ⚠️
	// msg1:= <-ch1
	// fmt.Println("✅ Received from ch1:",msg1)

	// msg2:= <-ch2
	// fmt.Println("☑️ Received from ch2:",msg2)

	//! SOLUTION 🚀
	select {
	case msg:= <-ch1:
	fmt.Println("✅ Received from ch1:",msg)

	case msg:= <-ch2:
	fmt.Println("☑️ Received from ch2:",msg)

	default:
	fmt.Println("🔴 Default Case - No channels ready! ")
	}

	//fmt.Println("End of the program.. 🔴")
}