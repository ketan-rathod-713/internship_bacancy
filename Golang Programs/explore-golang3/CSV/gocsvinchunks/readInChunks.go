package gocsvinchunks

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Customer struct {
	CustomerId       string `csv:"Customer Id"`
	FirstName        string `csv:"First Name"`
	LastName         string `csv:"Last Name"`
	Company          string `csv:"Company"`
	City             string `csv:"City"`
	Country          string `csv:"Country"`
	Email            string `csv:"Email"`
	SubscriptionDate string `csv:"Subscription Date"`
}

func ReadCsvFileInChunks() {
	// read in chunks of 10 mbs or some size
	// time.Sleep to check it

	file, err := os.Open("data/customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	var buffer = make([]byte, 1024)

	// Read in chunks of 10 mb
	for {
		n, err := file.Read(buffer)

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		// file ended
		if err == io.EOF {
			break
		}

		// read till n number hence write it
		fmt.Println(string(buffer[:n]))

		// Now unmarshal it but what if i want to read line by line is it possible as some records will be lost
		// 	var customers []Customer
		// 	err = gocsv.UnmarshalBytes(buffer[:n], &customers)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	fmt.Println()
		// 	fmt.Println()
		// 	fmt.Println()
		// 	fmt.Pr	intln()
		// 	fmt.Println(customers)

		// ! How to do it in chunks
		time.Sleep(time.Second * 15)

	}

	// above will give error because it is not reading csv data in line by line formate
	// Hence use csvReader for that purpose

	defer file.Close()
}

/*
How to read csv file

if file is small := reading line by line would be efficient as less memory requirement
if file is large and memory there : reading using buffer will reduce io operations

*/
