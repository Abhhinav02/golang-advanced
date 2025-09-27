package main

import (
	"fmt"
	"time"
)

func main() {
	// Sync between 2 channels 🚀
	numGoroutines:=3
	done:=make(chan int,3)

	for i:=range numGoroutines{
	go func(id int) {
		fmt.Printf("Goroutine %d working..\n",id)
		time.Sleep(time.Second)
		done<-id
	}(i)
	}

	for range numGoroutines{
		<-done // Wait for each goroutine to finish
	}

	fmt.Println("All goroutines are finished.. ✅")
}