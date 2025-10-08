package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	i:=0;
	for range ticker.C{
		i++
		fmt.Println(i)
	}

	// O.p
	// $ go run .
	// 1
	// 2
	// 3
	// 4
	// 5
	// ... So on...

}