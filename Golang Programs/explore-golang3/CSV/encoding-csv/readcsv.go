package encodingcsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsv() {
	file, err := os.Open("data/data.csv")
	if err != nil {
		log.Println(err)
	}

	// at the end of programm close file
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // there can be variable number of fields per record

	// ek hi bar me read karlo
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
	}

	// log.Println(records)

	// Print the CSV data
	for _, row := range records {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}
