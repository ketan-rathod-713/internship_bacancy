package generics2

import (
	"fmt"
)

// type Num2 int

// type Num interface {
// 	int | float64
// }

// Try 1
type Num interface {
	int
}

// Try 2
// type Num int

func Add(a int, b int) {
	fmt.Println("simple add", a, b)
}

func main() {
	var a, b int
	a = 10
	b = 20
	Add(a, b)
}
