package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	rune := 'a'
	fmt.Printf("rune with value %v of type %T\n", rune, rune)

	// to print the unicode representation of the code %c is used

	fmt.Printf("%c\n", rune)

	// to print the unicode code point represented by number %U is used
	fmt.Printf("%U\n", rune)

	// string is a slice of bytes

	// even though it is slice of bytes using range keyword willl iterate over stirng runes.

	myString := "❗hello"
	for index, char := range myString {
		fmt.Printf("Index: %d\tCharacter: %c\t Decimal:%v \tCode Point: %U\n", index, char, char, char)
	}
	// ohkk so some runes will be of size 1 byte will some can be of size 3 too. how the computer stores this information about it.
	// Since runes can be stored as 1, 2, 3, or 4 bytes, the length of a string may not always equal the number of characters in the string. Use the builtin len function to get the length of a string in bytes and the utf8.RuneCountInString function to get the number of runes in a string:

	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))
	// here so don't only focused on the string length ha ha
	// we have to be focused on runes not the bytes

	replacingInStringOfRunes()
}

func replacingInStringOfRunes() {
	str := "This is string"

	for index, char := range str {
		if char == 'i' {
			str = str[:index] + "o" + str[index+1:]
		}
	}
	fmt.Println(str)

	oldRune := '❗'
	newRune := 'I'
	log := "❗ recommended product ❗"
	for index, rune := range log {
		if rune == oldRune {
			log = log[:index] + string(newRune) + log[index+utf8.RuneLen(oldRune):] // index + 1 ki jagah size of old rune hona chahiye
		}
	}
	fmt.Println(log)
}

// NOT WORKING FOR REPLACING RUNES
// 	for index, rune := range log {
// 	if rune == oldRune {
// 		log = log[:index] + string(newRune) + log[index+1:]
// 	}
// }
