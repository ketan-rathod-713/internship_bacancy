package gocsvpackage

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func CsvToMaps() {
	file, err := os.Open("data/gocsvcreated.csv")
	if err != nil {
		log.Fatal(err)
	}

	// now csvtomaps
	mapsData, err := gocsv.CSVToMaps(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(mapsData)

	for _, data := range mapsData {
		log.Println(data["name"])
	}
}
