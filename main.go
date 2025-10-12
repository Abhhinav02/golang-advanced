package main

import (
	"fmt"
	"io"
	"os/exec"
)

// process spawning: interaction between process using Pipe()
func main() {
	pr,pw:=io.Pipe()
	cmd := exec.Command("grep","foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("food is good\nbar\nbaz\n"))
	}()

	output,err:=cmd.Output()

	if err!=nil{
		fmt.Println("⚠️ ERROR:",err)
		return
	}

	fmt.Println("✅ Output:",string(output))
	
}

//OP: ✅ Output: food is good

