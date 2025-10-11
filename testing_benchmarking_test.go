package main

import (
	"fmt"
	"testing"
)

//💡 benchmarking - used to measure the perfornance of the code. Especially, how long does it take to perform a function or program.

// we write benchmarking f(x) with names starting with 'benchmark'
// off. docs - https://pkg.go.dev/testing
func Add(a,b int)int{
	return a+b
}

func TestAddSubtests(t *testing.T){
	tests:= []struct{a,b,expected int}{
		{2,3,5},
		{0,0,0},
		{-1,1,0},
	}

	for _,test:=range tests{
		t.Run(fmt.Sprintf("Add(%d,%d)", test.a, test.b), func(t *testing.T){
			result:=Add(test.a, test.b)
			if result!=test.expected{
				t.Errorf("Result = %d; Expected = %d",result,test.expected)
			}
		})
	}

}

// O/P:
// $ go test testing_benchmarking_test.go
// ok      command-line-arguments  0.343s
