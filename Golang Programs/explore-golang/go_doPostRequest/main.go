package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformPostRequest()
}

func PerformPostRequest() {
	const url = "https://jsonplaceholder.typicode.com/posts"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename" : "Lets go with golang",
		"price":0,
		"platform": " Learn code online"
	}`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contentBytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(contentBytes))
}
