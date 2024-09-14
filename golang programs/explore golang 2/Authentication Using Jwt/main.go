package main

import (
	"auth/auth"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/home", auth.HomeHandler)
	http.HandleFunc("/refresh", auth.RefreshHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// everything should be done from client side
// refresh token should also be requested by client
// client need to send token everytime request
