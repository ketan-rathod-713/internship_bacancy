package example

import (
	"fmt"
)

type ErrNegativeSqrt float64 // ye namki error ham return karenge jab negative number aayenge

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return "An Error Occured"
	}
	return "No Error"
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func Example() {
	fmt.Println(ErrNegativeSqrt(-2).Error())
	fmt.Println(ErrNegativeSqrt(2).Error())
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
