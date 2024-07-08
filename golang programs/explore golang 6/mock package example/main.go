package main

import (
	"fmt"
	"math/rand"
)

// 1. Take non deterministic function calls and wrap them in interface
type randomNumberGenerator interface {
	randomInt(max int) int
}

// Now lets implement this interface using our standard librar rand struct
type standardRand struct{}

func (s standardRand) randomInt(max int) int {
	return rand.Intn(max)
}

// Now that we have one interface, now we can use it in our code like
// func divByRand(numerator int) int {
// 	return numerator / int(rand.Intn(10))
// }

func divByRand(numerator int, r randomNumberGenerator) int {
	return numerator / r.randomInt(10)
}

func main() {
	sr := standardRand{}
	ans := divByRand(20, sr)

	fmt.Println(ans)
}

// Now in our production call we would call this method by
// divByRandom(200, standardRand{})

// For testing write an implementation of the interface that uses testify mock
