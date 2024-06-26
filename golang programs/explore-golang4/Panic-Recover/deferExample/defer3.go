package deferExample

import (
	"fmt"
	"time"
)

func DeferClosure() {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 123 789
	}(a)

	a = 789

	time.Sleep(1 * time.Second)

	a = 200

	time.Sleep(2 * time.Second)
}
