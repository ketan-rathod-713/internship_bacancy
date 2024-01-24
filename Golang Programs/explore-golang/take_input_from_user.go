package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Using bufio package")
	// readInputUsingBufioReader()
	// readInputUsingFmtScan()
	readInputUsingFmtScanln()

	io.WriteString(os.Stdout, "wow great")
	// log.Fatal("error")
}

func readInputUsingFmtScanln(){ // why not working
	var text1 string
	fmt.Scanln(&text1)
	fmt.Println("the input text is ", text1)
}

func readInputUsingFmtScan(){
	var text1 string
	var text2 string
	fmt.Scan(&text1, &text2) // jitni bhi strings hogi vo successively ek ek argument me store hoti jaegi
	fmt.Println("the input text is ", text1, text2)
}

func readInputUsingBufioReader() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for pizza")

	// comma ok syntax // no try catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}

func readInputUsingBufIo() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		fmt.Println(text)
	}
}
