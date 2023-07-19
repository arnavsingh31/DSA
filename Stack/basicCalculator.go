package main

type ComputeStack struct {
	Items []int
	Top   int
}

func (cs *ComputeStack) Push(val int) {
	cs.Items = append(cs.Items, val)
	cs.Top++
}

func (cs *ComputeStack) Pop() int {
	topVal := cs.Items[cs.Top]
	cs.Items = cs.Items[:len(cs.Items)-1]
	cs.Top--
	return topVal
}

func calculate(s string) int {
	stack := &ComputeStack{
		Items: make([]int, 0),
		Top:   -1,
	}
	curr, res, sign := 0, 0, 1

	for _, r := range s {
		switch r {
		case ' ':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			curr = curr*10 + int(r-'0')
		case '+':
			res += curr * sign
			curr, sign = 0, 1
		case '-':
			res += curr * sign
			curr, sign = 0, -1
		case '(':
			stack.Push(res)
			stack.Push(sign)
			res, sign = 0, 1
		case ')':
			res += curr * sign
			res *= stack.Pop()
			res += stack.Pop()
			curr = 0
		}
	}

	return res + curr*sign
}
