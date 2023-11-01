package stack

/*
	LC #901
	T.C--> O(n)
	S.C--->O(n)
*/

type StockSpanPair struct {
	Price int
	Span  int
}

type StockSpanner struct {
	stack []StockSpanPair
}

func StockConstructor() StockSpanner {
	return StockSpanner{
		stack: make([]StockSpanPair, 0),
	}
}

func (s *StockSpanner) Next(price int) int {
	count := 1
	for len(s.stack) > 0 && s.stack[len(s.stack)-1].Price <= price {
		count += s.stack[len(s.stack)-1].Span
		s.stack = s.stack[:len(s.stack)-1]
	}

	s.stack = append(s.stack, StockSpanPair{Price: price, Span: count})

	return count
}

/*
	inputs := 100, 80, 60, 70, 60, 75, 85
	for 70 input :
		count = 1 + stackTopSpan = 1+1= 2
	for 75 input:
		count = 1 + stackTopSpan = 1+1 = 2,
		pop stack
		count = 2 + stackTopSpan = 2 + 2 = 4
		pop stack
	for 85 input:
		count = 1 + stackTopSpan= 1+4 = 5
		pop stack
		count = 5 + stackTopSpan = 5+1 = 6
		pop stack
	stack:- [(100,1),(85,6)  ]

*/
