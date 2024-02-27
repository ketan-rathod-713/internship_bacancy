package jsontwocsv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// use package json2csv

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsMarried bool   `json:"isMarried"`
}

func JsonToCsv() {
	// read json data
	file, err := os.Open("data/input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content []byte = make([]byte, 0)
	var buffer []byte = make([]byte, 100)

	for !errors.Is(err, io.EOF) {
		var n int = 0
		n, err = file.Read(buffer)

		fmt.Println(n)
		// time.Sleep(2 * time.Second)
		content = append(content, buffer[:n]...)
	}

	// fmt.Println(content)
	// got all data from file
	var data []Person
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	// Now got data convert it to csv

	output, err := os.Create("data/output.csv")
	defer output.Close()

	csvWriter := csv.NewWriter(output)
	for _, person := range data {
		var record []string
		record = append(record, fmt.Sprintf("%d", person.Id))
		record = append(record, person.Name)
		record = append(record,fmt.Sprintf("%v", person.Age) )
		record = append(record, fmt.Sprintf("%v",person.IsMarried))
		err = csvWriter.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}

	csvWriter.Flush()
}
