package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./data/students.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	reader.FieldsPerRecord = 3

	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			fmt.Println("End of file")
			break
		}

		if err != nil {
			fmt.Println("ERROR reading record", err)
			break
		}

		offset := reader.InputOffset()
		fmt.Println("Input offset", offset)

		line, column := reader.FieldPos(1)
		fmt.Println("Field Position Inside csv document ", line, column)

		fmt.Println(record)
	}

	// writting to a file

}
