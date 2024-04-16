// several type assertions in series

// cases here defines the types

// and those values are compared with the value hold by the given interface value

// switch v := i.(type) {
// case T:
//     // here v has type T
// case S:
//     // here v has type S
// default:
//     // no match; here v has the same type as i
// }

package typeassertion

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func TypeAssertion2() {
	do(21)
	do("hello")
	do(true)
}
