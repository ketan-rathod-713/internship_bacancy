package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("status code of response is : ", response.StatusCode)
	fmt.Println("Content Length", response.ContentLength)

	var responseString strings.Builder
	// content is in byte formate
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content) // when we have data in byte we can use it
	fmt.Println("byte count is ", byteCount)
	fmt.Println(responseString.String()) // whatever data is holding we get it.

	// beginners would prefer this but ig above one is great. We get much more control over string.
	// fmt.Println(string(content))
}