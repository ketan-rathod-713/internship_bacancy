package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Bufio implements buffered io
// it wraps io.Reader and io.Writter and creating another object Reader and Writter that also implements interface but provides buffering and some help for textual io.

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello This is a string reader"))

	// reader.ReadString() // Number of bytes that can be read from current buffer
	fmt.Println(reader.Buffered())

	// skips the next n bytes and returns the nummber of bytes discarded
	skipped, err := reader.Discard(4)
	fmt.Println("Discarded: ", skipped, err)

	// Peek method for peeking without advancing the reader
	peeked, err := reader.Peek(10)
	fmt.Println("Peeked:", string(peeked), err)

	// ReadByte reads sinlgle byte if no bytes then error
	// same Readbytes

	// b := 's'
	reader.ReadString('s')

	// const a = "30"

	const a = 4000
	f(a)
	f2(a)
	var m map[int]int
	// m[4]= 12
	val, ok := m[12321312321]
	fmt.Println(val, ok)
	fmt.Println(m)
	// fmt.Println(byte(rune(34)))
}

func f(a float64) {
	fmt.Println(a)
	return
}

func f2(a int) {
	fmt.Println(a)
	return
}
