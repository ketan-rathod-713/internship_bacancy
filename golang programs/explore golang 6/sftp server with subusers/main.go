package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"

	nomadApi "github.com/hashicorp/nomad/api"

	_ "github.com/lib/pq" // PostgreSQL driver
	"golang.org/x/crypto/bcrypt"
)

func createUser(username, password, folder string) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Store the hashed password in the database (example with PostgreSQL)
	db, err := sql.Open("postgres", "postgres://root:rootpass@localhost:5432/iceline-hosting?sslmode=disable")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	fmt.Println("Connected to database")

	_, err = db.Exec("INSERT INTO sftp_users (username, password, folder) VALUES ($1, $2, $3)", username, string(hashedPassword), folder)
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %w", err)
	}

	fmt.Println("Creating new user")
	// Create the user in the SFTP container
	cmd := exec.Command("docker", "exec", "sftp-server-container", "bash", "-c", fmt.Sprintf("useradd -m %s && echo '%s:%s' | chpasswd", username, username, password))
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to create user in Docker: %w - output: %s", err, string(output))
	}

	fmt.Println("Setting user permissions")
	// Create the user's directory and set permissions
	cmd = exec.Command("docker", "exec", "sftp-server-container", "bash", "-c", fmt.Sprintf("mkdir -p /home/%s/upload && chown -R %s:%s /home/%s/upload", username, username, username, username))
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to set user permissions in Docker: %w - output: %s", err, string(output))
	}

	return nil
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Folder   string `json:"folder"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := createUser(req.Username, req.Password, req.Folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s created successfully", req.Username)
}

func main() {
	http.HandleFunc("/create-user", createUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type nomadClient struct {
	client *nomadApi.Client
}

func NewNomadClient(nomadURL string) (*nomadClient, error) {
	nc, err := nomadApi.NewClient(&nomadApi.Config{
		Address:   nomadURL,
		TLSConfig: &nomadApi.TLSConfig{Insecure: true},
		SecretID:  os.Getenv("ACL_TOKEN"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Nomad client: %w", err)
	}

	// check if client is working
	_, _, err = nc.Jobs().List(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list jobs: %w", err)
	}

	return &nomadClient{
		client: nc,
	}, nil
}

func (n *nomadClient) getAllocations(ctx context.Context, jobID, namespace string) ([]*nomadApi.AllocationListStub, error) {
	allocs, _, err := n.client.Jobs().Allocations(jobID, false, &nomadApi.QueryOptions{Namespace: namespace})
	if err != nil {
		return nil, fmt.Errorf("failed to get allocations: %w", err)
	}

	if len(allocs) == 0 {
		return nil, errors.New("no allocations found")
	}

	// allocations are returned in sorted order
	sort.Slice(allocs, func(i, j int) bool {
		return allocs[i].CreateTime > allocs[j].CreateTime
	})

	return allocs, nil
}

func (n *nomadClient) RunCommand(ctx context.Context, jobID, namespace string, stdin io.Reader, stdout, stderr io.Writer, cmd string, args ...string) (int, error) {
	allocs, err := n.getAllocations(ctx, jobID, namespace)
	if err != nil {
		return 0, err
	}

	var allocationId string
	for _, a := range allocs {
		if a.TaskGroup == jobID {
			allocationId = a.ID
			break
		}
	}

	if allocationId == "" {
		return 0, errors.New("no suitable allocation found")
	}

	alloc, _, err := n.client.Allocations().Info(allocationId, &nomadApi.QueryOptions{Namespace: namespace})
	if err != nil {
		return 0, fmt.Errorf("failed to get allocation info: %w", err)
	}

	var termSizeCh chan nomadApi.TerminalSize

	command := append([]string{cmd}, args...)
	log.Println("Executing command:", strings.Join(command, " "))
	exitCode, err := n.client.Allocations().Exec(ctx, alloc, jobID, true, command, stdin, stdout, stderr, termSizeCh, &nomadApi.QueryOptions{Namespace: namespace})
	if err != nil {
		return exitCode, fmt.Errorf("failed to execute command: %w", err)
	}

	return exitCode, nil
}
