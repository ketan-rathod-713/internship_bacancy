package main

import (
	"fmt"
	"reflect"
)

func main() {
	// empty mape /make(map[key-type]val-type).
	mp := make(map[string]int)

	// DOUBT -> can we use new keyword here to create the map. See it.

	mp["k1"] = 32
	mp["k2"] = 45

	fmt.Println(mp["k1"])

	// delete one key value pair
	delete(mp, "k1")

	// remove all key value pairs
	clear(mp)

	// value and isPresent
	_, isPresent := mp["k2"]
	fmt.Println("prs:", isPresent)

	// in same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n, reflect.TypeOf(n))
}
