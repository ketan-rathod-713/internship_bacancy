package gocsvpackage

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// struct fields will map to column names in csv file
// csv tag provided for smooth mapping

// we can use csv:"-" to ignore that field
type Person struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	Gender string `csv:"gender"`
}

// marshal this slice of structs into a CSV format using the gocsv.MarshalFile() function, which takes a pointer to the slice of structs and the file object to write the CSV data to
var data []Person = []Person{
	Person{Name: "Ketan", Age: 20, Gender: "male"},
	Person{Name: "Aman", Age: 20, Gender: "male"},
	Person{Name: "Tridip", Age: 20, Gender: "male"},
	Person{Name: "Asmita", Age: 20, Gender: "female"},
	Person{Name: "Bhumika", Age: 20, Gender: "female"},
}

func Marshaling() {
	// create csv file to write
	file, err := os.Create("data/gocsvcreated.csv")
	if err != nil {
		log.Fatal(err)
	}

	// close file at the end of the program
	defer file.Close()

	// Now marshal our given data
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("data saved in gocsvcreated.csv")
}
