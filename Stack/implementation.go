package main

type Stack struct {
	Items []interface{}
	Top   int
}

func Constructor() *Stack {
	return &Stack{
		Items: make([]interface{}, 0),
		Top:   -1,
	}
}

func (s *Stack) Push(val interface{}) {
	s.Items = append(s.Items, val)
	s.Top++
}

func (s *Stack) Pop() interface{} {
	if s.Top == -1 {
		return 0
	}
	topVal := s.Items[s.Top]
	s.Items = s.Items[:len(s.Items)-1]
	s.Top--

	return topVal
}

func (s *Stack) Peek() interface{} {
	if s.Top == -1 {
		return 0
	}

	return s.Items[s.Top]
}

// func main() {
// 	stack := Constructor()
// 	stack.Push("A")
// 	stack.Push("B")
// 	stack.Push("C")
// 	stack.Push("D")
// 	log.Printf("stack consists of %v", stack.Items)
// 	log.Print(stack.Peek())
// 	log.Print(stack.Pop())
// 	log.Printf("stack consists of %v", stack.Items)
// 	log.Print(stack.Peek())
// }
