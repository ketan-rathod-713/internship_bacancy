package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

const url = "https://lco.dev"

// Make sure to close the connection. It is your responsibility to close the connection.

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if(err != nil){
		panic(err)
	} else {
		fmt.Println("Type of response is ", reflect.TypeOf(response))
		fmt.Println("Response is ", response)
	}

	// Callers responsibility to close the connection
	defer response.Body.Close()

	// we can read the response // majority read done by ioutil

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println("\n\nThe content is ", content)
}
