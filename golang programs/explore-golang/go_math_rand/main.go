package main

import (
	"fmt"
	"math"
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	d := rand.Intn(6) + 1
	return d
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	d := rand.Float64() * 12
	return d
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	slice := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

func main() {
	fmt.Println("Roll a die", RollADie())
	fmt.Println(GenerateWandEnergy())
	fmt.Println("shuffle animals", ShuffleAnimals())

	fmt.Println(math.Pow(20, 3))
}

// TODO: IMP rand package and shuffle array
