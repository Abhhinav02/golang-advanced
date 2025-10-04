package main

import "fmt"

func producer(ch chan <-int){
	for i:=range 5{
		ch<-i
	}
	close(ch)
}

func filter(in <- chan int, out chan <-int){
	for val:=range in{
		if val%2==0{
			// pass to the sender/other channel
			out<-val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1,ch2)

	for val:=range ch2{
		fmt.Println("Val:",val)
	}
// OP:
// $ go run .
// Val: 0
// Val: 2
// Val: 4
}