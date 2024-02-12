package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/hello2" {
		http.NotFound(w, r)
		return // here we have to return also ha ha
	}

	if r.Method != "POST" { // Only Allowing Post Request on this route.
		http.Error(w, "Method is not supported.", http.StatusNotFound) // 3rd parameter is status code of response writter
		return
	}

	w.Write([]byte("hii ha ha"))
}

func main() {
	// We’ll use the HandleFunc function to add route handlers to the web server.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/hello2", hello2Handler)

	log.Fatal(http.ListenAndServe(":8080", nil)) // Here nil  because we are setting up http2 here hence no need to define it.
}
