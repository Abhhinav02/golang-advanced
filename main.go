package main

import (
	"fmt"
	"time"
)

func periodicTask(){
	fmt.Println("⌛ Performing periodic task at:",time.Now())
}

func main() {
	// Periodic Task exec()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop() // Turn off ticker

	for{
	select{
	case <-ticker.C:
		periodicTask()
	}
	}

	

// O.p
// $ go run .
// 	⌛ Performing periodic task at: 2025-10-08 16:28:24.8883307 +0530 IST m=+6.001331601
// ⌛ Performing periodic task at: 2025-10-08 16:28:25.8885089 +0530 IST m=+7.001509801
// ⌛ Performing periodic task at: 2025-10-08 16:28:26.8879589 +0530 IST m=+8.000959801
// ⌛ Performing periodic task at: 2025-10-08 16:28:27.8880838 +0530 IST m=+9.001084701
// ⌛ Performing periodic task at: 2025-10-08 16:28:28.8878862 +0530 IST m=+10.000887101
// ⌛ Performing periodic task at: 2025-10-08 16:28:29.8882013 +0530 IST m=+11.001202201
// .... so on...
	

}