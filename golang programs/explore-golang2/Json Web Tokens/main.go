package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secreatKey = "SECREAT"

func main() {
	myHandler := http.HandlerFunc(homeHandler)
	http.Handle("/", authMiddleware(myHandler))

	token, err := generateJWT()
	if err != nil {
		log.Println("ERROR OCCURED", err)
	}

	fmt.Println(token)

	http.ListenAndServe(":8080", nil)
}

func authMiddleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before request handling")
		originalHandler.ServeHTTP(w, r)
		fmt.Println("After request handling")
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Handling")
	w.Write([]byte("Hello World"))
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString("wow great")
	return tokenString, err
}
