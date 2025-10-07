package main

import (
	"context"
	"fmt"
	"time"
)

// context.TODO() example
func checkEvenOrOdd(ctx context.Context, num int )string{
	select {
	case <-ctx.Done():
		return "Operation cancelled!"
	default:
		if num%2==0{
			return  fmt.Sprintf("%d is even ✅",num)
		} else{
			return  fmt.Sprintf("%d is odd ☑️",num)
		}	
	}
	
}

func main() {
	ctx:= context.TODO()
	result:= checkEvenOrOdd(ctx,5)
	fmt.Println("Result with context.TODO() :",result)

	ctx = context.Background()

	ctx,cancel:= context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// once again
	result = checkEvenOrOdd(ctx, 10)
	fmt.Println("Result with timeout context:",result)
	 defer cancel()

	

	// now, sleep..
	time.Sleep(2*time.Second)
	result = checkEvenOrOdd(ctx, 15)
	fmt.Println("Result with after sleep/timeout:",result)

	//O/P:
	// $ go run .
	// Result with context.TODO() : 5 is odd ☑️
	// Result with timeout context: 10 is even ✅
	// Result with after sleep/timeout: Operation cancelled!
	

}
