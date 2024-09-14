package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"

	nomadApi "github.com/hashicorp/nomad/api"
)

type nomadClient struct {
	client *nomadApi.Client
}

func NewNomadClient(nomadURL string, aclToken string) (*nomadClient, error) {
	nc, err := nomadApi.NewClient(&nomadApi.Config{
		Address:   nomadURL,
		TLSConfig: &nomadApi.TLSConfig{Insecure: true},
		SecretID:  aclToken, // PUT ACL TOKEN HERE
	})
	if err != nil {
		return nil, err
	}

	// check if client is working
	_, _, err = nc.Jobs().List(nil)
	if err != nil {
		return nil, err
	}

	return &nomadClient{
		client: nc,
	}, nil
}

func main() {
	fmt.Println("Hello, World!")

	jobId := "6b05ca2d-c762-4986-8279-8cea4f36f4a6-ca8b3b11-58f9-4493-8ee6-b90dda88efb2-Minecraft-79679fed-5416-4b22-85dd-da6be741ee2a"
	namespace := jobId
	taskName := ""
	taskGroupName := ""

	log.Println("data", namespace, taskName, taskGroupName)

	c, err := NewNomadClient("https://29ed-103-156-142-118.ngrok-free.app", "cc36b192-072d-fe63-56d8-f8fdd7cf282d")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// allocationsList, err := c.getAllocations(ctx, jobId, namespace)
	// if err != nil {
	// 	log.Fatal("can not get allocations", err)
	// }

	// log.Println("allocations got", allocationsList)

	exitcode, err := c.RunCommand(ctx, jobId, namespace, os.Stdin, os.Stdout, os.Stdout, "/bin/sh", "mkdir new")
	if err != nil {
		log.Fatal("error running command", err)
	}

	log.Println("exitcode", exitcode)
}

func (n *nomadClient) RunCommand(ctx context.Context, jobID, namespace string, stdin io.Reader, stdout, stderr io.Writer, cmd string, args ...string) (int, error) {
	allocs, err := n.getAllocations(ctx, jobID, namespace)
	if err != nil {
		return 0, err
	}
	log.Println("allocations got", allocs)

	// iterate all allocations and get it's id by comparing job id
	for _, a := range allocs {
		fmt.Println("allocation id", a.ID)
		fmt.Println("allocation job id", a.JobID)
		fmt.Println("allocation name:", a.Name)
		fmt.Println("task group", a.TaskGroup) // it will return Server_Installation
		fmt.Println("task states", a.TaskStates)
		// how allocation name is given, if i can rename the allocation name at the time of registering then yeah i can do it. just by comparing the allocation name.
	}

	var allocationId string = ""
	for _, a := range allocs {
		if a.TaskGroup == "Server_Installation" {
			fmt.Println("server installation allocation present")
		}

		if a.TaskGroup == jobID {
			fmt.Println("minecraft allocation present")
			allocationId = a.ID
		}
	}
	fmt.Println("allocation id of minecraft allocation", allocationId)

	alloc, _, err := n.client.Allocations().Info(allocs[1].ID, &nomadApi.QueryOptions{Namespace: namespace})
	if err != nil {
		log.Println("error in getting allocation", err)
		return 0, err
	}
	log.Println("Allocations", alloc.ID)

	var termSizeCh chan nomadApi.TerminalSize

	command := []string{cmd}
	command = append(command, args...)
	log.Println(command)
	exitCode, err := n.client.Allocations().Exec(ctx, alloc, jobID,
		true, command, stdin, stdout, stderr, termSizeCh, &nomadApi.QueryOptions{Namespace: namespace})
	if err != nil {
		log.Println("error in executing command", err)
		return 0, err
	}
	log.Println(alloc, "****", jobID)
	return exitCode, err
}

func (n *nomadClient) getAllocations(ctx context.Context, jobID, namespace string) ([]*nomadApi.AllocationListStub, error) {
	allocs, _, err := n.client.Jobs().Allocations(jobID, false, &nomadApi.QueryOptions{Namespace: namespace})
	if err != nil {
		return nil, err
	}

	if len(allocs) == 0 {
		return nil, errors.New("no allocations")
	}
	for i := 0; i < len(allocs); i++ {
		log.Println(allocs[i].Name)

	}
	sort.Slice(allocs, func(i, j int) bool {
		return allocs[i].CreateTime > allocs[j].CreateTime
	})

	return allocs, nil
}
