package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime) // current time is given

	// we will use this standard time for formatting
	// This is given in documentation hence see it whenever need formate of date
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Month is type of time.Month and others are integers
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
