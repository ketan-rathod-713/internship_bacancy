// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.

package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return factorial(n-1) * n
}

func main() {
	ans := factorial(3)
	fmt.Printf("factorial of %d is %d", 3, ans)
}
