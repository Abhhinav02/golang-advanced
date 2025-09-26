package main

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main-thread and run in the BG and they return to join the main-thread once the functions are finished/ready to return any value.

// Goroutines do not stop the program-flow and are non-blocking. (kinda async-await in JS)

func main() {

	var err error

   fmt.Println("Before sayHello().. ✅")
   go sayHello()
   fmt.Println("After sayHello().. ☑️")


   // workaround f(x) as we cannot set goroutine directly on -> err = doWork() 
   go func(){
	err = doWork()
   }() // ✅

   // err = go doWork() ❌
   
   if err!=nil{
	fmt.Println("🔴ERROR: ",err)
   }else{
	fmt.Println("Work completed successfully ✔️")
   }

   go printNums()
   go printLetters()



   time.Sleep(2*time.Second)

// ERROR ⚠️ - cannot use it after the sleep-timer❌
//     if err!=nil{
// 	fmt.Println("🔴ERROR: ",err)
//    }else{
// 	fmt.Println("Work completed successfully ✔️")
//    }

}

func sayHello() {
	time.Sleep(1 * time.Second) // waits for 1 second
	fmt.Println("Hello from Goroutine")
}

// Other examples
func printNums(){
	for i:=range 5{
		fmt.Println("🔵 Number: ",i,time.Now())
		time.Sleep(100 *time.Millisecond)
	}
}

func printLetters(){
	for _,letter:= range "abcdef"{
		fmt.Println("🟠 Letter: ",string(letter),time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

// concurrency vs parallelism
func doWork() error{
	// simulate work
	time.Sleep(1 * time.Second)

	// simulate err.
	return  fmt.Errorf("🔴 An ERROR occured in doWork()!")
}

