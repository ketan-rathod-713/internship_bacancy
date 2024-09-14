// The interface type that specifies zero methods is known as the empty interface:

// interface{}
// An empty interface may hold values of any type. (Every type implements at least zero methods.)

// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

package example

import "fmt"

type MyString string

type MyStruct struct {
	something int
}

func (s *MyStruct) Print() {
	fmt.Println(s)
}

type P interface {
	Print()
}

func DescribeEmptyInterface() {
	fmt.Println("I can put any repeated value on variable of type interface")
	fmt.Println("as it is not datatype hence interface me koi bhi value dal sakte he")

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = MyString("hello") // main.MyString type ha ha
	describe(i)

	i = MyStruct{something: 42}
	describe(i)

	var p P
	p = &MyStruct{something: 42}
	p.Print()
}

// It can take any value as every thing implements 0 methods on it
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
