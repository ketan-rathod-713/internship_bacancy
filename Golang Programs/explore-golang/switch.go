package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
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
