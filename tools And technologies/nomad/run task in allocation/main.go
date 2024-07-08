package main

import (
	"context"
	"io"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/hashicorp/nomad/api"
)

// CustomIO struct that implements io.Reader and io.Writer
type CustomIO struct {
	data []byte
	mu   sync.Mutex
}

// NewCustomIo creates a new instance of CustomIO
func NewCustomIo() *CustomIO {
	return &CustomIO{
		data: make([]byte, 0, 2048),
	}
}

// Read implements the io.Reader interface
func (c *CustomIO) Read(p []byte) (n int, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	log.Println("Started reading")

	if len(c.data) == 0 {
		return 0, io.EOF
	}

	n = copy(p, c.data)
	c.data = c.data[n:] // consume the data read

	log.Println("Read:", string(p[:n]))
	return n, nil
}

// Write implements the io.Writer interface
func (c *CustomIO) Write(p []byte) (n int, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data = append(c.data, p...)
	n = len(p)

	log.Println("Write:", string(p))
	return n, nil
}
func main() {

	// helper := NewCustomIo()

	// helper.Write([]byte("good"))

	// var p []byte = make([]byte, 1024)
	// n, err := helper.Read(p)

	// if err != nil {
	// 	log.Println("err reading ")
	// }

	// log.Println(n, string(p))

	// Initialize a new Nomad API client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("failed to create Nomad client: %v", err)
	}
	var jobId string = "e0aafbc8-d408-4362-981e-9a310d28a656-2912f838-196b-4c77-a717-7311343e530d-2-minecraft-f0a6e961-a71f-48c2-a095-09b1cea85761"
	var namespace string = jobId
	allocs, _, err := client.Jobs().Allocations(jobId, true, &api.QueryOptions{
		Namespace: namespace,
	})

	if err != nil {
		log.Println("enable to get allocations")
	}

	if len(allocs) == 0 {
		panic("no allocations")
	}

	// sort allocations according to the time they are created
	sort.Slice(allocs, func(i, j int) bool {
		return allocs[i].CreateTime > allocs[j].CreateTime
	})

	log.Println("Allocations found are ", allocs)

	// Now we have got our allocation // lets play with it to run a command

	// Info is used to retrieve a single allocation if we know the allocation id

	allocation, _, err := client.Allocations().Info(allocs[0].ID, &api.QueryOptions{Namespace: namespace})

	if err != nil {
		log.Panic("error in getting allocation", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initialize io.Reader, io.Writer, and io.Writer
	var stdin io.Reader = os.Stdin
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stdout
	var command []string = make([]string, 0)
	command = append(command, "/bin/sh")
	log.Println("Command we are executing is", command)

	// issue command
	// go func() {
	// 	time.Sleep(100 * time.Millisecond)
	// 	stdout.Write([]byte("ls\n"))

	// }()

	var termSizeCh chan api.TerminalSize

	exitCode, err := client.Allocations().Exec(ctx, allocation, jobId, true, command, stdin, stdout, stderr, termSizeCh, &api.QueryOptions{Namespace: namespace})

	// log.Println("Data available inside helper", string(helper.data))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Process exited with exit code", exitCode)
}
