package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.gofile.io/servers")
	if err != nil {
		fmt.Println("error ", err)
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading bytes", err)
	}

	fmt.Println(string(bodyBytes))
}
