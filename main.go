package main

import (
	"fmt"
	"time"
)

// real-world app. example
type ticketRequest struct{
	personId int
	numOfTickets int
	cost int
}

// simulate processing of ticketRequests by creating workers
func ticketProcessor(requests <-chan ticketRequest, results chan<-int){
	for req := range requests{
		fmt.Printf("Processing %d ticket(s) of personId %d with total cost of %d\n", req.numOfTickets, req.personId, req.cost)

		// simulate ticket-processing time/delay
		time.Sleep(time.Second)
		results <- req.personId

	}
}

func main() {
	numOfReqs := 5
	price:=5
	ticketRequests:= make(chan ticketRequest, numOfReqs)
	ticketResults:= make(chan int)

	// Create workers/ticket-processor
	for range 3{
		go ticketProcessor(ticketRequests,ticketResults)
	}

	// Send ticket-requests
	for i:= range numOfReqs{
		ticketRequests <-ticketRequest{personId: i+1, numOfTickets:(i+1)*2, cost:(i+1)*price}
	}

	close(ticketRequests)

	for range numOfReqs{
		fmt.Printf("\n🟢Ticket for personId %d processed successfully",<-ticketResults)
	}
	
}
// O/P
// $ go run .
// Processing 2 ticket(s) of personId 1 with total cost of 5
// Processing 4 ticket(s) of personId 2 with total cost of 10
// Processing 6 ticket(s) of personId 3 with total cost of 15
// Processing 8 ticket(s) of personId 4 with total cost of 20

// 🟢Ticket for personId 3 processed successfully
// 🟢Ticket for personId 1 processed successfully
// 🟢Ticket for personId 2 processed successfullyProcessing 10 ticket(s) of personId 5 with total cost of 25

// 🟢Ticket for personId 4 processed successfully
// 🟢Ticket for personId 5 processed successfully