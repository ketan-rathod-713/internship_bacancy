package main

import "fmt"

func main() {
	days := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	fmt.Println(days)

	// normal for loop
	for d := 0; d < len(days); d++ { // there is no ++d ha ha
		fmt.Print(days[d], " ")
	}

	// using range
	for key, value := range days {
		fmt.Println("days[", key, "]:", value)
	}

	// just like while loop
	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 1 {
			goto lco
		}

		if rogueValue == 3 {
			rogueValue++ // so that next we skip the one value ha ha
			continue
		}

		if rogueValue == 5 {
			break
		}

		fmt.Println(rogueValue)
		rogueValue++
	}

	// label was declared here and we need to use it.
	// should i place it at the end of the file always.
lco:
	fmt.Println("Jumping at here ha ha")
	fmt.Println("once again")
}
