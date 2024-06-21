package main

import "fmt"

func main() {

	// fmt.Println("Using Append Method")

	// strByte := []byte{}

	// strByte = fmt.Append(strByte, "Hello world")

	// strByte = fmt.Append(strByte, "wow it is looking good ha ha")

	// fmt.Println(string(strByte))

	// // Appendf same but for formatted string
	// // Appendln new line appended

	// fmt.Println("")
	// fmt.Println("Using Errorf method")
	// const name, id = "bueller", 17
	// err := fmt.Errorf("user %q (id %d) not found", name, id)
	// fmt.Println(err.Error())

	// // Fscan and FPrint
	// // both read and writes to io.Writter and io.Reader

	fmt.Println("Reading ")

	// Sscan and sprintf is used when we want to return thatt string to a variable

	// Formatting float values
	value := 300
	fmt.Printf("%[1]d %[2]d", 20, value)

	fmt.Printf("%d %d \n", 20, 40)
	// TODO: DOUBT why it was printing % when new line is not used.

	var s string
	var i int
	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
	fmt.Println(s, i)

	var str string
	var intI int
	fmt.Println("")
	fmt.Sscanf("9237293423good45", "%14s%d", &str, &intI)
	fmt.Println(str, intI)

}

// Stringer interface
// String() string method to be implemented.
