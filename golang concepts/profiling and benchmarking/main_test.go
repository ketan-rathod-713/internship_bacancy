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
// For eg.
// go test -bench=^BenchmarkMemoryWaste -benchmem -benchtime=1s -count=10 > refactor.bench

// now it is time to compare the results of those two benchmarks
// we have tool -> benchstat to compare the results

// to get benchstat :- go get golang.org/x/perf/cmd/benchstat
// Then compile it by :- go install golang.org/x/perf/cmd/benchstat
// Now we are ready to use it.

// For eg.
// benchstat origin.bench refactor.bench
// it will output
/**
goos: linux
goarch: amd64
pkg: profiling_example
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
              │ origin.bench │           refactor.bench            │
              │    sec/op    │   sec/op     vs base                │
MemoryWaste-8    465.4n ± 2%   163.4n ± 4%  -64.89% (p=0.000 n=10)

              │ origin.bench │           refactor.bench           │
              │     B/op     │    B/op     vs base                │
MemoryWaste-8    2040.0 ± 0%   896.0 ± 0%  -56.08% (p=0.000 n=10)

              │ origin.bench │           refactor.bench           │
              │  allocs/op   │ allocs/op   vs base                │
MemoryWaste-8     8.000 ± 0%   1.000 ± 0%  -87.50% (p=0.000 n=10)
**/

// The output is divided into three parts :-
// Execution Time: Reduced by 64.89%.
// Memory Usage: Reduced by 56.08%.
// Memory Allocations: Reduced by 87.50%.

// allocations decreased from 8 to 1.
// n stands for number of times we ran the benchmark.
// p stands for probability that the result is significant. it should be less then 0.05

// There is also a thing that is benchmark in parallel.
// some times we want to test functions that are running in parrallel.

func BenchmarkMemoryWasteParallel(b *testing.B) {
	// GOMAXPROCS

	// spawns as many as goroutines as the numbers of GOMAXPROCS.
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemoryWaste(100)
		}
	})
}

// output wills say that it is not better to test different things
// concurrent and sequential

// Profiling in go :-
// While benchmarking gives you a comparative view, profiling is about in depth understanding of your software.

// benchmarking gives external view while profiling gives internal view.
// profiling is a next step after benchmarking.

// we can use pprof tool for profiling purpose.
// it is tool for visualization and analysis of Go programs' memory profiles.

// In golang there are different types of profiles :-

/**
- CPU Profile
	-> Common profile used. used for identifying the cpu intensive tasks in your go code.
- Memory Profile
	- In use space
		-> live objects
	- Allocated space
		-> all the objects that has been allocated even if they garbage collected
- Block Profile
	-> measures the time your goroutines waiting for channel or like
- Mutex Profile
	-> how often goroutines are blocked to acquire mutex lock
- Goroutine Profile
	-> provides a snapshot of all current goroutines, used for diagnostic deadlocks, goroutines leaks or other concurrency issues.
- Threadcreate Profile
	-> number of operating system threads TODO:
- Debug Traces
	->
**/
