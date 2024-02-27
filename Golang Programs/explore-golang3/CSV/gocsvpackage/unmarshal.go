package gocsvpackage

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func Unmarshal() {
	file, err := os.Open("data/gocsvcreated.csv")
	if err != nil {
		panic(err)
	}

	var data []Person
	gocsv.Unmarshal(file, &data)
	log.Println(data)
}
