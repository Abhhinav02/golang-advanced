package main

import (
	"fmt"
)

 func main() {
	ch := make(chan int)

	go func(){
		ch <-1
		close(ch)
	}()

	for {
		select{
		case msg,ok:= <-ch:
			if !ok{
				fmt.Println("Channel closed 🔴")
				// cleanup activities.. 🧹
				return
			}
			fmt.Println("Received:",msg)
		}
	}
 }	



//  func main() {
// 	ch := make(chan int)

// 	go func(){
// 		time.Sleep(2* time.Second) 
// 		ch <-1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("✅Received from ch:",msg)
// 	case <- time.After(3* time.Second):
// 		fmt.Println("⏳ Timeout.. ")
// 	}
//  }


// multiplexing - kinda like switch-case
// func main() {
// 	ch1:= make(chan int)
// 	ch2:= make(chan int)

// 	go func(){
// 		time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()

// 	go func(){
// 		time.Sleep(time.Second)
// 		ch2 <- 2
// 	}()


// 	time.Sleep(2* time.Second) // Waiting..

// 	for range 2{
// 	select {
// 	case msg:= <-ch1:
// 	fmt.Println("✅Received from ch1:",msg)

// 	case msg:= <-ch2:
// 	fmt.Println("☑️Received from ch2:",msg)

// 	default:
// 	fmt.Println("🔴 Default Case - No channels ready! ")
// 	}
// }


// 	fmt.Println("###### End of the program ######")
// }