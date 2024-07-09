package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/nomad/api"
)

// defining custom input/output struct

type IO struct {
	readBufferLock sync.Mutex
	readBuffer     []byte // for the logic purpose
	outputBuffer   []byte // for summarizing final output from the command
	cancelFunc     context.CancelFunc
}

func NewIO(cancelFunc context.CancelFunc) *IO {
	return &IO{
		readBuffer: make([]byte, 0, 2048),
		cancelFunc: cancelFunc,
	}
}

// read method will be used by the nomad to read the command we sent to it.
func (i *IO) Read(p []byte) (int, error) {
	// time.Sleep(2 * time.Second)
	i.readBufferLock.Lock()
	defer i.readBufferLock.Unlock()

	n := copy(p, i.readBuffer)
	i.readBuffer = i.readBuffer[n:]

	return n, nil
}

// write method will be used by nomad to write the output of the command we sent
func (i *IO) Write(p []byte) (int, error) {
	// write or send message to anyone and return the length of the data sent
	// fmt.Printf("%v", string(p))

	// NOTE: i can not cancel the context here because the nomad server is doing multiple writes for the same thing.

	// task completed hence
	// fmt.Println("cancelling context")
	// i.cancelFunc()

	i.outputBuffer = append(i.outputBuffer, p...)

	// Alternate Solution
	// Check that if it has 2 hashtags 1 at start and 1 at end then done ha ha

	return len(p), nil
}

// read a single command from the user
func (i *IO) readCommand(command string) {
	i.readBufferLock.Lock()
	defer i.readBufferLock.Unlock()

	i.readBuffer = append(i.readBuffer, fmt.Sprintf("%s\n", command)...)
}

func main() {
	// Initialize a new Nomad API client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("failed to create Nomad client: %v", err)
	}

	// We have same jobId and Namespace for our project
	var jobId string = "4a1cb83b-1c37-482c-82f9-36bd045164f3-a721cb72-7401-4db4-bcd4-480e5e9da69b-2-minecraft-b0c42615-456f-46e5-9142-c3c421fdea86"
	var namespace string = jobId
	allocs, _, err := client.Jobs().Allocations(jobId, true, &api.QueryOptions{
		Namespace: namespace,
	})

	if err != nil {
		log.Println("unable to get allocations:", err)
		return
	}

	if len(allocs) == 0 {
		panic("no allocations found")
	}

	// sort allocations according to the time they are created
	sort.Slice(allocs, func(i, j int) bool {
		return allocs[i].CreateTime > allocs[j].CreateTime
	})

	log.Println("Allocations found are", allocs)

	allocation, _, err := client.Allocations().Info(allocs[0].ID, &api.QueryOptions{Namespace: namespace})

	if err != nil {
		log.Panic("error in getting allocation:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// initialize custom io
	io := NewIO(cancel)

	stdin := io
	stdout := io
	stderr := io

	io.readCommand("ls -ltr")

	command := []string{"/bin/sh"}
	log.Println("Command we are executing is", command)

	var termSizeCh chan api.TerminalSize

	exitCode, err := client.Allocations().Exec(ctx, allocation, jobId, true, command, stdin, stdout, stderr, termSizeCh, &api.QueryOptions{Namespace: namespace})

	if err != nil {
		log.Println("error", err)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Printing io output buffer \n", formattedOutput(string(io.outputBuffer)))
	fmt.Println("")

	log.Println("Process exited with exit code", exitCode)
}

// gets the string between two # hashtags.
func formattedOutput(s string) string {
	start := strings.Index(s, "#")
	end := strings.LastIndex(s, "#")

	if start == -1 || end == -1 || start == end {
		return "" // No valid substring found
	}

	return s[start+1 : end]
}
