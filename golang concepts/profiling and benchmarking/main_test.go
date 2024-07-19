package main

import "testing"

// Benchmark and testing are same overall as both using same testing package
// When doing benchmarking, we use Benchmark prefix before the function name

func BenchmarkMemoryWaste(b *testing.B) {

	// benchmark donesn't run for once, it runs b.N times
	for i := 0; i < b.N; i++ {
		MemoryWaste(100)
	}
}

// To run the benchmark, use `go test -bench=^BenchmarkMemoryWaste`
// we can improve above by using -benchmem flag to measure memory usage
// as we can see for now we have 8 allocs/op

// for now output looks like this
/**
goos: linux
goarch: amd64
pkg: profiling_example
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkMemoryWaste-8           2626154               469.7 ns/op          2040 B/op          8 allocs/op
PASS
ok      profiling_example       1.699s
**/

// we can improve above thing by using -benchtime flag to limit time for each benchmark and with -count flag to run benchmark multiple times

// For eg.

/**
➜  profiling and benchmarking git:(main) ✗ go test -bench=^BenchmarkMemoryWaste -benchmem -benchtime=1s -count=10
goos: linux
goarch: amd64
pkg: profiling_example
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkMemoryWaste-8           2627367               451.4 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2438770               452.8 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2645301               448.6 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2645377               458.8 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2596620               458.2 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2619052               454.4 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2640777               454.8 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2673088               453.7 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2596560               453.0 ns/op          2040 B/op          8 allocs/op
BenchmarkMemoryWaste-8           2649500               470.3 ns/op          2040 B/op          8 allocs/op
PASS
ok      profiling_example       16.570s
**/

// why benchmarking more then once ?
// because when we run benchmark multiple times, we can get more accurate and reliable results because our system can do warm up.

// we can store our benchmark results in the file using > operator
// For eg.
// go test -bench=^BenchmarkMemoryWaste -benchmem -benchtime=1s -count=10 > origin.bench
// output of benchmark will get stored in origin.bench file

// now lets improve our function to use less allocations and memory requirements, then we will come to benchmark again.
