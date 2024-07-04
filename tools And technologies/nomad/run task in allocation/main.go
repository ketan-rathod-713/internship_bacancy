package main

import (
	"context"
	"io"
	"log"
	"os"
	"sort"
	"time"

	"github.com/hashicorp/nomad/api"
)

func main() {
	// Initialize a new Nomad API client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("failed to create Nomad client: %v", err)
	}
	var jobId string = "ddd10034-3597-4c4a-9256-5bcaeb3b78a6-893cc98f-4ac8-482a-acea-f6398459c18b-2-minecraft-fc32714e-977e-4b4e-baea-174ed2d540c5"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Initialize io.Reader, io.Writer, and io.Writer
	var stdin io.Reader = os.Stdin
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr
	var command []string = make([]string, 0)
	command = append(command, "/bin/sh")
	log.Println("Command we are executing is", command)

	var termSizeCh chan api.TerminalSize

	exitCode, err := client.Allocations().Exec(ctx, allocation, jobId, true, command, stdin, stdout, stderr, termSizeCh, &api.QueryOptions{Namespace: namespace})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Process exited with exit code", exitCode)
}
