package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	for range 2 {
		select {
		case <-timer1.C:
			fmt.Println("Timer 1 expired!")
		case <-timer2.C:
			fmt.Println("Timer 2 expired!")
		}
	}
	// O/P:
	// $ go run .
	// Timer 1 expired!
	// Timer 2 expired!
}
