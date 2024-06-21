package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hashicorp/nomad/api"
)

func main() {
	// Define flags
	listJobs := flag.Bool("listjobs", false, "List all jobs")
	flag.Parse()

	// Create a new Nomad client
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Nomad client: %v", err)
	}

	log.Println(client.Address())
	// Execute CLI commands based on flags
	if *listJobs {
		log.Println("Listing all jobs")
		listAllJobs(client)
	}

	createJob(client, "golang_job")
}

func createJob(client *api.Client, jobId string) {
	log.Println("Creating a job")

	res, meta, err := client.Jobs().Register(&api.Job{ID: &jobId, }, &api.WriteOptions{Region: ""})

	log.Println(res, meta)
	log.Println(err)
}

func listAllJobs(client *api.Client) {
	jobs, _, err := client.Jobs().List(nil)
	if err != nil {
		log.Fatalf("Failed to list jobs: %v", err)
	}

	for _, job := range jobs {
		fmt.Printf("Job ID: %s, Job Name: %s\n", job.ID, job.Name)
	}
}
