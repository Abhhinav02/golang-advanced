package main

import (
	"fmt"
	"time"
)

//💡 TASK/DIY: Multiple tickers

func main() {

	ticker1 := time.NewTicker(time.Second)
	ticker2 := time.NewTicker(500 * time.Millisecond)

	stop:= time.After(5*time.Second)
	defer ticker1.Stop()
	defer ticker2.Stop()

	for{
	select{
	case tick1:= <-ticker1.C:
		fmt.Println("🔵 1st Tick at:",tick1)
	case tick2:= <-ticker2.C:
		fmt.Println("🟢 2nd Tick at:",tick2)
	// ⚠️ using time.After() not recommended inside long loops (ok, for this example)
	case <- stop:
		fmt.Println("Stopping both the tickers.. ☑️")	
		return
		}
	}
	
// O.p:
// $ go run .
// 🟢 2nd Tick at: 2025-10-08 16:59:27.8078304 +0530 IST m=+0.500000001
// 🔵 1st Tick at: 2025-10-08 16:59:28.3078304 +0530 IST m=+1.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:28.3078304 +0530 IST m=+1.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:28.8078304 +0530 IST m=+1.500000001
// 🔵 1st Tick at: 2025-10-08 16:59:29.3078304 +0530 IST m=+2.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:29.3078304 +0530 IST m=+2.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:29.8078304 +0530 IST m=+2.500000001
// 🔵 1st Tick at: 2025-10-08 16:59:30.3078304 +0530 IST m=+3.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:30.3078304 +0530 IST m=+3.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:30.8078304 +0530 IST m=+3.500000001
// 🔵 1st Tick at: 2025-10-08 16:59:31.3078304 +0530 IST m=+4.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:31.3078304 +0530 IST m=+4.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:31.8078304 +0530 IST m=+4.500000001
// 🔵 1st Tick at: 2025-10-08 16:59:32.3078304 +0530 IST m=+5.000000001
// 🟢 2nd Tick at: 2025-10-08 16:59:32.3078304 +0530 IST m=+5.000000001
// Stopping both the tickers.. ☑️
}