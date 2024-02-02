package main

import "fmt"

func main() {
	stack := new(Stack)
	stack.top = -1
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	fmt.Println(stack.Top())
}

type Stack struct {
	top  int
	data []int
}

func (s *Stack) Push(x int) {
	s.top++
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int{
	s.top--
	return s.data[s.top + 1]
}

func (s *Stack) Top() int {
	s.top--
	return s.data[s.top+1]
}
