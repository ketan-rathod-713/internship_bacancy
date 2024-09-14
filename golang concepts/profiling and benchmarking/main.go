package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime/pprof"
	"strings"
	"sync"
	"time"

	_ "net/http/pprof"
	// TODO: we have to install it in order to use the pprof

	"github.com/PuerkitoBio/goquery"
)

func MemoryWaste(n int) []int {
	var data []int = make([]int, 0, n)

	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	return data
}

var leakyStorage = make(map[int]*goquery.Document)
var mu = sync.Mutex{}

func getRandomString() string {
	// get random string
	randomString := make([]byte, 500)
	for i := 0; i < 500; i++ {
		randomString[i] = byte(rand.Intn(26) + 'a')
	}
	return string(randomString)
}

func generateHtml() string {
	var html string
	html += "<html>\n<body>\n"
	for i := 0; i < 10000; i++ {
		html += fmt.Sprintf("<div>%d</div>\n", i)
	}

	html += "</body>\n</html>"

	return html
}

func main() {
	fmt.Println("started profiling...")
	var cpuprofile = flag.String("cpuprofile", "cpu.pprof", "write cpu profile to `file`")
	var memoryprofile = flag.String("memoryprofile", "memory.pprof", "write cpu profile to `file`")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not create CPU profile: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "could not start CPU profile: %v\n", err)
			os.Exit(1)
		}
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 100; i++ {
		fmt.Println("creating document", i)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(generateHtml()))

		if err != nil {
			fmt.Println("failed to parse document")
			return
		}

		// Simulate memory leak by storing the parsed documents
		mu.Lock()
		leakyStorage[time.Now().Nanosecond()] = doc
		mu.Unlock()
	}

	file, err := os.Create(*memoryprofile)
	if err != nil {
		fmt.Println("unable to create memory profile file")
		return
	}

	err = pprof.WriteHeapProfile(file)
	if err != nil {
		fmt.Println("unable to write heap profile")
		return
	}

	fmt.Println("profiling completed")

	// we can use pprof to debug the http server also,
	http.ListenAndServe(":8080", nil)

	// by using
	// go tool pprof http://localhost:8080/debug/pprof/allocs?debug=1

}

// using pprof we can do profiling
// for that generate the pprof file by `go run main.go cpuprofile=cpu.pprof`

// For eg.
// go tool pprof cpu.pprof

// important commands
// top 10 - show top cpu intensive operations
// gcDrain -> job for the garbage collector
// we can also see pdf by typing pdf command

//
