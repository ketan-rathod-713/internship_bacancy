package main

import (
	"fmt"
	"io"
	"os"
)

//

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	fmt.Println(file)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1)

	for {
		// read a chunk
		n, err := file.Read(buf) // at the end of the file returns 0
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break // end of line EOF
		}

		// Printing or writting chunk
		fmt.Println(string(buf))

		// For writting to file don't write whole buffer only write till n byttes from buffer ha ha
		// REFER https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
	}

	// Read Whole File at once
	bytes, err := os.ReadFile("file.txt")
	fmt.Println(string(bytes))

	os.WriteFile("out.txt", []byte("wow great"), 0644)
}
