package main

import "fmt"

type Pair struct {
	Index int
	Temp  int
}

type Stack struct {
	data []Pair
	Top  int
}

func (s *Stack) Push(val Pair) {
	s.Top++
	s.data = append(s.data, val)
}

func (s *Stack) Pop() Pair {
	if s.Top == -1 {
		return Pair{Index: -1}
	}
	i := s.data[s.Top]
	s.data = s.data[:s.Top]
	s.Top--
	return i
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans)
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack Stack = Stack{Top: -1}
	for i := 0; i < len(temperatures); i++ {
		fmt.Println(i, ans)
		if stack.Top == -1 { // empty then directly push
			// do nothing then
		} else { // stack is not empty
			for stack.Top != -1 && stack.data[stack.Top].Temp < temperatures[i] {
				// evaluate for them
				pr := stack.Pop()
				ans[pr.Index] = i - pr.Index // current index - pehle ka index
			}
		}
		// at last push current one
		pair := Pair{i, temperatures[i]}
		stack.Push(pair)
	}

	return ans
}
