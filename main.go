package main

import (
	"context"
	"fmt"
)

// context - instance of a struct / obj.
func main() {
	todoCtxt := context.TODO()
	bgCtxt := context.Background()

	ctxt:= context.WithValue(todoCtxt, "name","Skyy")
	ctxtBg:= context.WithValue(bgCtxt, "city","Kolkata")

	fmt.Println(ctxt)
	fmt.Println(ctxt.Value("name"))

	fmt.Println(ctxtBg)
	fmt.Println(ctxtBg.Value("city"))
}