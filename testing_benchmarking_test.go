package main

import (
	"math/rand"
	"testing"
)

// profiling (bonus info.) - provides detailed insights to the performance of our app., including CPU usage, memory allocation and goroutine activities

func GenerateRandSlice(size int)[]int{
	slice:=make([]int, size)

	for i:= range slice{
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int)int{
	sum:=0

	for _,v:=range slice{
		sum+=v
	}
	return sum
}

// Now, let's test

func TestGenerateRandSlice(t *testing.T){
	size:=100
	slice:= GenerateRandSlice(size)

	// assertion
	if len(slice) != size{
		t.Errorf("Expected slice size: %d; Received: %d",size, len(slice))
	}
}

// benchmark the randSlice f(x)
func BenchmarkGenerateRandSlice(b *testing.B){
	for b.Loop(){
		GenerateRandSlice(1000)
	}
}

// benchamrk the sumOfSlice f(x)
func BenchmarkSumSlice(b *testing.B){
	slice:= GenerateRandSlice(1000)
	/*
	ResetTimer zeroes the elapsed benchmark time and memory allocation counters and deletes user-reported metrics. It does not affect whether the timer is running. More accurate.
	*/
	b.ResetTimer()

	for b.Loop(){
		SumSlice(slice)
	}
}


// 💡O/P:
// $ go test -bench=. -memprofile mem.pprof testing_benchmarking_test.go|grep -v 'cpu'
// goos: windows
// goarch: amd64
// BenchmarkGenerateRandSlice-8       75289             16818 ns/op
// BenchmarkSumSlice-8              2554164               452.3 ns/op
// PASS
// ok      command-line-arguments  2.913s

// 💡 Another way (INTERACTIVE MODE 🟢):
// $ go tool pprof mem.pprof
// File: main.test.exe
// Build ID: C:\Users\ASUS\AppData\Local\Temp\go-build3061445528\b001\main.test.exe2025-10-12 01:29:34.8496867 +0530 IST
// Type: alloc_space
// Time: 2025-10-12 01:29:37 IST
// Entering interactive mode (type "help" for commands, "o" for options)
// (pprof) top
// Showing nodes accounting for 624.86MB, 99.76% of 626.36MB total
// Dropped 17 nodes (cum <= 3.13MB)
//       flat  flat%   sum%        cum   cum%
//   624.86MB 99.76% 99.76%   624.86MB 99.76%  command-line-arguments.GenerateRandSlice
//          0     0% 99.76%   624.86MB 99.76%  command-line-arguments.BenchmarkGenerateRandSlice
//          0     0% 99.76%   624.86MB 99.76%  testing.(*B).run1.func1    
//          0     0% 99.76%   624.86MB 99.76%  testing.(*B).runN
// (pprof) list GenerateRandSlice
// Total: 626.36MB
// ROUTINE ======================== command-line-arguments.BenchmarkGenerateRandSlice in C:\Users\ASUS\Desktop\golang-advanced\testing_benchmarking_test.go
//          0   624.86MB (flat, cum) 99.76% of Total
//          .          .     41:func BenchmarkGenerateRandSlice(b *testing.B){
//          .          .     42:   for b.Loop(){
//          .   624.86MB     43:           GenerateRandSlice(1000)        
//          .          .     44:   }
//          .          .     45:}
//          .          .     46:
//          .          .     47:// benchamrk the sumOfSlice f(x)
//          .          .     48:func BenchmarkSumSlice(b *testing.B){     
// ROUTINE ======================== command-line-arguments.GenerateRandSlice in C:\Users\ASUS\Desktop\golang-advanced\testing_benchmarking_test.go
//   624.86MB   624.86MB (flat, cum) 99.76% of Total
//          .          .     10:func GenerateRandSlice(size int)[]int{    
//   624.86MB   624.86MB     11:   slice:=make([]int, size)
//          .          .     12:
//          .          .     13:   for i:= range slice{
//          .          .     14:           slice[i] = rand.Intn(100)      
//          .          .     15:   }
//          .          .     16:   return slice
// (pprof)

// 🟢 And many more....