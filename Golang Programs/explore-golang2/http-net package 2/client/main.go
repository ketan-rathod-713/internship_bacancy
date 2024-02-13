package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {

	/* Making Get Request To Server */
	res, err := http.Get("http://localhost:8080?name=ketan,aman&student=true")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.ContentLength)
	log.Println(res.Status)
	log.Println(res.StatusCode)
	log.Println(res.Header)

	cb, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(cb))

	/* Making Post Request */

	data := `{name:"ketan"}`

	res, err = http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer([]byte(data)))

	if err != nil {
		log.Fatal(err)
	}

	cb, _ = io.ReadAll(res.Body)
	log.Println(string(cb))

	/* Sending Post Request Over Urlencoded Data */
	res, err = http.Post("http://localhost:8080/", "xxx-urlencoded-form", bytes.NewBuffer([]byte(data)))

	if err != nil {
		log.Fatal(err)
	}

	cb, _ = io.ReadAll(res.Body)
	log.Println(string(cb))
}
