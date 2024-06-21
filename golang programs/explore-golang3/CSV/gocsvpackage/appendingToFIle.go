package gocsvpackage

import (
	"encoding/csv"
	"log"
	"os"
)

// Open file with append tag so that we don't overwrite previous data in our file
func AppendingToFile() {
	// Most users will be using os.Open
	// THis is generalised function which requires flags to be passed

	// IMPORTANT // TODO
	// If you're using os.O_APPEND flag alone, without specifying any other flags, such as os.O_WRONLY or os.O_CREATE, it might not work as expected. The os.O_APPEND flag by itself is intended to be used in conjunction with other flags, typically os.O_WRONLY for write-only access or os.O_RDWR for read-write access.

	file, err := os.OpenFile("data/gocsvcreated.csv", os.O_APPEND | os.O_RDWR, 0644) // 0644 are permissions
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// create new writter
	writter := csv.NewWriter(file)
	defer writter.Flush()

	row := []string{"David", "30", "Male"}
	err = writter.Write(row)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Appended data in file")
}
