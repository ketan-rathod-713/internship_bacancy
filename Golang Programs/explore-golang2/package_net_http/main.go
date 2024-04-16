package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"sync"
	// t "crypto/tls" // working with crypto/tls package
)

type countHandler struct {
	mu sync.Mutex
	n  int
}

func (c *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
	fmt.Fprintf(w, "count is %d", c.n)
}

func main() {
	// t.Server()

	// tls.ExampleTLS()

	fmt.Println("Canonical header key", http.CanonicalHeaderKey("content-type"))

	var data = `{"name": "ketan"}`
	fmt.Println(http.DetectContentType([]byte(data)))

	// no writes should be done as Error does not close the connection.
	// it returns the plain text
	// http.Error(nil, "error", 202)

	http.Handle("/", new(countHandler))

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./data/file.txt")
	})

	// ! what to do here.
	// TODO:
	// http.HandleFunc("/file2", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFileFS(w, r, )
	// })

	

	// listens on tcp connection and calls serve method http.Serve(listener, handler)
	fmt.Println("Serve started on port 8080")
	http.ListenAndServe(":8080", nil)

	// alternate :
	// ListenAndServeTLS
	// it expects https conenctions
	//
}
