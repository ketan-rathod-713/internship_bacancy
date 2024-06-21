package main

import (
	"fmt"

	"example.com/greetings/greetings"
	"example.com/hello/package1"
)

func main() {
	str := package1.Greeting1()
	fmt.Println(str)

	greetings.Hello("great")
}
