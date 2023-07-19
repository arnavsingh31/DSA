package main

import (
	"strconv"
)

type RPNStack struct {
	Items []int
	Top   int
}

func (s *RPNStack) Push(a int) {
	s.Items = append(s.Items, a)
	s.Top++
}

func (s *RPNStack) Pop() int {
	topVal := s.Items[s.Top]
	s.Items = s.Items[:len(s.Items)-1]
	s.Top--
	return topVal
}

// LC #150
func evalRPN(tokens []string) int {
	arithematicOperatorMap := map[string]func(int, int) int{
		"+": func(i1, i2 int) int { return i1 + i2 },
		"-": func(i1, i2 int) int { return i1 - i2 },
		"*": func(i1, i2 int) int { return i1 * i2 },
		"/": func(i1, i2 int) int { return i1 / i2 },
	}

	stack := &RPNStack{
		Items: make([]int, 0),
		Top:   -1,
	}

	for _, str := range tokens {
		if op, exist := arithematicOperatorMap[str]; !exist {
			num, _ := strconv.Atoi(str)
			stack.Push(num)
		} else {
			v1 := stack.Pop()
			v2 := stack.Pop()
			stack.Push(op(v2, v1))
		}
	}
	return stack.Pop()
}
