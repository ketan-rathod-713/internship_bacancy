package gocsvpackage

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func Unmarshalling() {
	file, err := os.Open("data/gocsvcreated.csv")
	if err != nil {
		log.Fatal(err)
	}

	// now unmarshal using gocsv
	var data []Person
	err = gocsv.UnmarshalFile(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range data {
		log.Println(record.Name, record.Age, record.Gender)
	}

	log.Println("successfully got data in struct field")
}
