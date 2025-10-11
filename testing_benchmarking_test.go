package main

import (
	"testing"
)

// MULTIPLE BENHCMARK F(X)

func Add(a,b int)int{
	return a+b
}

func BenchmarkSmallInput(b *testing.B){
	// for range b.N{
	// Loop returns true as long as the benchmark should continue running.
	for b.Loop(){
		Add(2,3)
	}
}

func BenchmarkMediumInput(b *testing.B){
	// for range b.N{
	// Loop returns true as long as the benchmark should continue running.
	for b.Loop(){
		Add(200,300)
	}
}

func BenchmarkLargeInput(b *testing.B){
	// for range b.N{
	// Loop returns true as long as the benchmark should continue running.
	for b.Loop(){
		Add(2000,3000)
	}
}

// O/P:
// $ go test -bench=. -benchmem  testing_benchmarking_test.go|grep 
// -v 'cpu:'
// goos: windows
// goarch: amd64
// BenchmarkSmallInput-8           864131830                1.358 ns/op           0 B/op          0 allocs/op
// BenchmarkMediumInput-8          783273967                1.384 ns/op           0 B/op          0 allocs/op
// BenchmarkLargeInput-8           801466148                1.352 ns/op           0 B/op          0 allocs/op
// PASS
// ok      command-line-arguments  3.740s
