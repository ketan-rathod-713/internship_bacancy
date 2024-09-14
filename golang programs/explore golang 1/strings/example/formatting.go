package example

import "fmt"

func Formatting() {
	n := 1.2

	s := fmt.Sprintf("%20.2f", n)
	fmt.Println(s)

	t := fmt.Sprintf("%020.2f", n)
	fmt.Println(t)
}
