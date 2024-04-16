package main

import (
	"fmt"
	"interfacesexample/example"
	"interfacesexample/stringer"
	typeassertion "interfacesexample/type_assertion"
)

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

	fmt.Println("Describe interface")

	example.DescribeEmptyInterface()

	fmt.Println("Interface Nil value will produce a run time error because it is not a data type such that it can produce a compile time error")

	example.InterfaceNilValue()

	fmt.Println("Interface shape example implementing area method")

	example.InterfaceExampleOfShapeCircle()

	fmt.Println("Interface stringer example")

	stringer.InterfaceStringerExample()
	stringer.InterfaceStringerExample2()
	stringer.InterfaceStringerExample3()

	fmt.Println("Type Assertion of interface")
	typeassertion.TypeAssertion1()
	typeassertion.TypeAssertion2()

	fmt.Println("Interface value good example")
	typeassertion.InterfaceValueExample()

	// TODO: IMP Type assertion of interface
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
