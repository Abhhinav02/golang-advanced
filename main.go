package main

import (
	"fmt"
	"time"
)

// tickers⏱️ + timers⌛
// Scheduling logging📝
// Polling for updates⚙️
// Handling ticker-stops gracefully🪶

func main() {

	ticker := time.NewTicker(time.Second)
	stop:= time.After(5*time.Second)
	defer ticker.Stop()

	for{
	select{
	case tick:= <-ticker.C:
		fmt.Println("⌚ Tick at:",tick)
	case <- stop:
		fmt.Println("Stopping ticker.. ☑️")	
		return
		}
	}
	
// O.p
// $ go run .
// ⌚ Tick at: 2025-10-08 16:44:58.8357751 +0530 IST m=+1.000557801
// ⌚ Tick at: 2025-10-08 16:44:59.8357751 +0530 IST m=+2.000557801
// ⌚ Tick at: 2025-10-08 16:45:00.8357751 +0530 IST m=+3.000557801
// ⌚ Tick at: 2025-10-08 16:45:01.8357751 +0530 IST m=+4.000557801
// ⌚ Tick at: 2025-10-08 16:45:02.8357751 +0530 IST m=+5.000557801
// Stopping ticker.. ☑️

}