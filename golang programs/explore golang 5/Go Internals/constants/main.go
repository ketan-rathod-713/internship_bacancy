package main

import (
	"fmt"
	"reflect"
)

// type of this is untyped string
// untyped constant is just a value, which is not given a type yet
const str = "this is string"

type MyString string

func main() {
	var s string = str
	fmt.Println(s, reflect.TypeOf(s))

	// it works
	var mystring MyString = str
	fmt.Println(mystring, reflect.TypeOf(mystring))

	// Below code will produce an error
	// var mystring2 MyString = s

	var mystring3 MyString = MyString(s)
	fmt.Println(mystring3)

	// ? our untyped string constant, it has the helpful property that, since it has no type, assigning it to a typed variable does not cause a type error. That is, we can write

	// ? as untyped contants do not have type so assigning them to any type does not give an error

	// ? These untyped string constants are strings, of course, so they can only be used where a string is allowed, but they do not have type string.

	// ? Default type
	// ? “if the constant is untyped, how does str get a type in this variable declaration?” The answer is that an untyped constant has a default type, an implicit type that it transfers to a value if a type is needed where none is provided. For untyped string constants, that default type is obviously string.

	// ? One way to think about untyped constants is that they live in a kind of ideal space of values, a space less restrictive than Go’s full type system. But to do anything with them, we need to assign them to variables, and when that happens the variable (not the constant itself) needs a type, and the constant can tell the variable what type it should have. In this example, str becomes a value of type string because the untyped string constant gives the declaration its default type, string.

	// Default type deternmined by syntax.

	// for untyped string only string type can be done
	// for numeric type there are variations like int, floating point constants to float64, rune constants to rune int32. and imaginary constats to complex128

	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", 'x', 'x')
	fmt.Printf("%T %v\n", 0i, 0i)


	// above rule also same for the untyped booleans
	type MyBool bool
    const True = true
    const TypedTrue bool = true
    var mb MyBool
    mb = true      // OK
    mb = True      // OK
    // mb = TypedTrue // Bad
    fmt.Println(mb)

	// ? Once a given a type boolean variables can not be mixed.

	// Floats

	// it is same as that of booleans
	// one thing is that there are two types of float
	// float32 and float64

	const Huge = 1e1000

	// we can declare such a large number, but we can not print it nor assign it.
	// hence constant number precision se likhe ja sakte he after all it is a value right.

	// ? same for complex numbers

	// ? Integers

	
}
