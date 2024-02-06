package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("T is nill ha ha")
		return
	}
	fmt.Println(t.S) // Nill ki value access karne ka try kar rahe he ham ha ha
	// go doesnt give us null pointer exception
}

func main() {
	var i I // interface

	// DOUBT // should it take only pass by reference value only
	var t *T
	i = t // passing nil value
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()



	k := T{S: "great"} // ye value return karta he pura block store hoga variable me
	fmt.Println(k)

	p := &T{"wow"}
	fmt.Println(*p)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
	// here type of interface will show a tuple of value and concrete type
}

// Note that an interface value that holds a nil concrete value is itself non-nil.


