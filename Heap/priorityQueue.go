package heap

type PriorityQueue []*Item

type Item struct {
	data int
	x, y int
}

func (h PriorityQueue) Len() int {
	return len(h)
}

func (h PriorityQueue) Less(i, j int) bool {
	return h[i].data < h[j].data
}

func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *PriorityQueue) Pop() interface{} {
	lastIndex := len(*h) - 1
	lastItem := (*h)[lastIndex]
	(*h)[lastIndex] = nil
	*h = (*h)[:lastIndex]
	return lastItem
}
