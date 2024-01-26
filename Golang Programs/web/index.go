package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/", madeFunction)

	// http.Handle("/", m)
	http.ListenAndServe(":8080", m)
}

func madeFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
