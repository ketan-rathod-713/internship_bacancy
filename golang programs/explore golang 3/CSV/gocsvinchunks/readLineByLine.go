package gocsvinchunks

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func ReadCsvLineByLine() {
	file, err := os.Open("data/customers.csv")

	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	var customers []Customer = make([]Customer, 0)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			log.Println("Read csv complete")
			break
		}

		// Now unmarshal record in my struct and process it

		// TODO : can i unmarhal it here.
		var customer Customer = Customer{
			FirstName: record[1],
		}
		customers = append(customers, customer)
		fmt.Println(customer)

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(2 * time.Second)
	}
}

func ReadLineByLineUsingCustomLogic() {
	file, err := os.Open("data/customers.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

}
