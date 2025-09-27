package main

import (
	"fmt"
	"time"
)

func main() {
	// Syncing DATA EXCHANGE 🚀
	// Here, we use channels to pass data
	// Using that data in our app

	numGoroutines:= 5
	data:=make(chan string)

	go func() {
		for i:=range numGoroutines{
			data <- "☑️ Hello "+fmt.Sprint(i)
			time.Sleep(100 * time.Millisecond)
	}
	close(data)
	}()
	

	for val:=range data{
	fmt.Println("Received value:",val, ": ",time.Now())
	} 	
}