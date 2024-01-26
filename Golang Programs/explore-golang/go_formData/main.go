package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

func main() {
	PerformPostFormRequest()
}

// Performing Post Request Form Request
// Why we need form type data and how this all happening
func PerformPostFormRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	// formdata
	data := url.Values{}
	data.Add("firstname", "ketan")
	data.Add("lastname", "Rathod")
	data.Add("email", "ketanrtd1@gmail.com")

	response, _ := http.PostForm(myurl, data)

	defer response.Body.Close()

	// why we need io util here
	// DOUBT
	// what is the type of response.Body
	contentBytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Type of Response.Body is ", reflect.TypeOf(response.Body))
	fmt.Println(string(contentBytes))

}
