package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// This is a good syntax. // sometimes value comes in web request and we just assign it and check.
	if num := 3; num < 10 { // initialization, comparison
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	var err string
	if err != "" {
		fmt.Println("wow")
	}

	Switch()
}

func Switch() {
	rand.Seed(time.Now().UnixNano()) // we have seeded so that we get random number every time
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice with value 1")
	case 2:
		fmt.Println("Dice with value 2")
	case 3:
		fmt.Println("Dice with value 3")
		fallthrough // if we want here to not have break. it should also go in case below it.
	case 4:
		fmt.Println("Dice with value 4")
		fallthrough
	default:
		fmt.Println("This is default for 5 and 6")
	}

	// Switch without condtion is same as that of if else struct

	j := 20
	switch {
	case j < 20:
		fmt.Println("less then 20")
	case j < 40:
		fmt.Println("less then 40")
	case j == 40:
		fmt.Println("Equal to 40")
	}
}
