package main

import (
	"fmt"
	"net/url"
	"reflect"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asdasd"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	// to get information from url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query() // all the query parameters in better representation
	fmt.Println("Type fo Query() is ", reflect.TypeOf(qparams))
	// This are of type of key value pairs

	fmt.Println(qparams["coursename"])

	// i can also use for loop and range here

	for _, val := range qparams {
		fmt.Println("Param is ", val)
	}

	// DOUBT
	// to construct the url
	partsOfUrl := &url.URL{ // we always have to pass the address here // why ??
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Url created is ", anotherUrl)
}
