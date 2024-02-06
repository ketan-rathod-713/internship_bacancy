package go_methods_interfaces

import "fmt"

type Vertex struct {
	something int
}

func (v Vertex) method1() { // here actual change is not happening we need to return it if we want
	v.something = 20
}

// Pointer Receivers
func (v *Vertex) method2() { // here actual change is not happening we need to return it if we want
	v.something = 20
}

func AnotherFunction(v Vertex, value int) {
	v.something = value
}

func AnotherFunctionPtr(v *Vertex, value int) {
	v.something = value
}

func main() {
	var v Vertex = Vertex{something: 12}
	v.method1() // this will not work //
	fmt.Println(v)

	v.method2()
	fmt.Println(v)

	AnotherFunction(v, 40)
	fmt.Println(v)

	AnotherFunctionPtr(&v, 50)
	fmt.Println(v)
}

// Note
// functions with a pointer argument must take a pointer:
// var v Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK
// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK

// Above same thing can be done in reverse direction too.
// if pass by value and pass address then methods will find its waay but functions will produce an error.

// DOUBT
// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)