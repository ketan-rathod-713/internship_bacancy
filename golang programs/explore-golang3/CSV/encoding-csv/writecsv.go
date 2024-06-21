package encodingcsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var dataToWrite [][]string = [][]string{
	{
		"name", "heading2",
	},
	{
		"ketan", "233",
	},
	{
		"aman", "435",
	},
	{
		"rahul", "213",
	},
}

func WriteToCsv() {
	// create csv file
	file1, err := os.Create("data/created.csv")
	if err != nil {
		log.Println(err)
	}

	// Close file at the end of this function
	defer file1.Close()
	defer fmt.Println("I am deffered")

	// Now create new csv writter so that we can write to that file
	writter := csv.NewWriter(file1)

	for _, data := range dataToWrite {
		writter.Write(data)
	}

	// writes are buffered hence to ensure that all data is written to io.Writter from buffer
	// we need to call flush method to flush all data in io.writter

	defer writter.Flush()
}
