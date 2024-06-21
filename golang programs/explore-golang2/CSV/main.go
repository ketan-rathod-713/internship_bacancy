package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Name       string `csv:"name"`
	Percentage string `csv:"percentage"`
	Class      int    `csv:"class"`
}

func main() {
	file, err := os.Open("data/students.csv")
	if err != nil {
		log.Println("ERROR", err)
	}

	defer file.Close()

	// var student Student
	// record, err := csv.NewReader(file).Read()
	// if err != nil {
	// 	log.Println("ERROR", err)
	// }

	// fmt.Println(record)

	// records, err := csv.NewReader(file).ReadAll()
	// if err != nil {
	// 	log.Println("ERROR", err)
	// }

	// fmt.Println(records)

	// csvFile, err := os.Create("employee.csv")
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }
	// defer csvFile.Close()

	// csvwriter := csv.NewWriter(csvFile)

	// employees := [][]string{{"aman", "great"}, {"jetan", "wow"}, {"wjat", "wonder"}}

	// for _, record := range employees {
	// 	_ = csvwriter.Write(record)
	// }
	// csvwriter.Flush()

	// if csvwriter.Error() != nil {
	// 	fmt.Println(csvwriter.Error())
	// }

	file3 := csv.NewReader(file)
	fmt.Println(file3.FieldsPerRecord)

	record, err := file3.Read()
	fmt.Println(len(record))
}
