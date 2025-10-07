package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ctxt with cancel() example
// manually run the cancel func()

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("🔴 Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working.. ✅")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ctx:= context.Background();
	ctx, cancelFx:=context.WithCancel(ctx)
	// defer cancelFx() ❌

	go func(){
		time.Sleep(2 * time.Second) // simulating heavy-task (time consuming op.)
		cancelFx() // manually/only after tast completion
		}()

	ctx = context.WithValue(ctx, "reqID", "abcd1234")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	requestId:= ctx.Value("reqID")

	if requestId!=nil{
		fmt.Println("Request ID:",requestId)
	}else{
		fmt.Println("No request ID found!")
	}

	logWithCtxt(ctx, "This a test logger message ☑️")
}

// create ctxt with value: ctx.withValue()
// extracrting the value: ctx.Value()

func logWithCtxt(ctx context.Context, message string){
reqIdVal:=ctx.Value("reqID")
log.Printf("ReqID: %v - %v",reqIdVal,message)
}

//OP-
//$ go run .
// Working.. ✅
// Working.. ✅
// Working.. ✅
// Working.. ✅
// 🔴 Work cancelled: context canceled
// Request ID: abcd1234
// 2025/10/08 05:25:28 ReqID: abcd1234 - This a test logger message ☑️ ....