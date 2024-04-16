package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	stack := new(Stack)
	stack.data = make([]string, 20)

	for {
		ShowOptions()
		var option string
		_, err := fmt.Scan(&option)

		if err != nil {
			fmt.Println("Err occured reading option", err)
		}

		option_Int, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println("Err occured reading option", err)
		}

		switch option_Int {
		case 1:
			fmt.Println("Push data to stack, type data you want to push...")
			var data string
			fmt.Scan(&data)
			err := stack.Push(data)

			if err != nil {
				fmt.Println("Error Pushing Data", err)
			}
		case 2:
			data, err := stack.Pop()
			if err != nil {
				fmt.Println("Error Popping Data", err)
			} else {
				fmt.Println("Poped data", data)
			}
		case 3:
			stack.data = make([]string, 20)
		case 4:
			fmt.Println(stack.data)
		}
	}
}

func ShowOptions() {
	fmt.Println("Apply any one operation on stack")
	fmt.Println("1. Push")
	fmt.Println("2. Pop")
	fmt.Println("3. reset")
	fmt.Println("4. print")
}

type Stack struct {
	top  int
	data []string
}

func (s *Stack) Push(x string) error {

	// if s.top+1 > len(s.data) {
	// 	return errors.New("Stack is full")
	// }

	s.top++
	// s.data = append(s.data, x)
	s.data[s.top] = x

	return nil
}

func (s *Stack) Pop() (string, error) {

	if s.top == 0 || s.top == -1 {
		return "", errors.New("Empty Stack")
	}

	s.top--
	return s.data[s.top], nil
}

func (s *Stack) Top() (string, error) {
	if s.top >= 0 {
		return s.data[s.top], nil
	}

	return "", errors.New("Stack is Empty")
}
