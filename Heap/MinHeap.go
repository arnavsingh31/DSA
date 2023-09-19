package heap

type CustomMinHeap []*Cell

func (mh *CustomMinHeap) push(c *Cell) {

	*mh = append(*mh, c)
	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent].val > (*mh)[index].val {
			mh.swap(parent, index)
			index = parent
		} else {
			return
		}
	}
}

func (mh *CustomMinHeap) pop() {
	lastIndex := len(*mh) - 1

	// swap
	mh.swap(0, lastIndex)

	// pop last
	*mh = (*mh)[:lastIndex]

	mh.Heapify(0, len(*mh)-1)
}

func (mh *CustomMinHeap) Heapify(index int, n int) {
	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left <= n && (*mh)[left].val < (*mh)[smallest].val {
		smallest = left
	}

	if right <= n && (*mh)[right].val < (*mh)[smallest].val {
		smallest = right
	}

	if smallest != index {
		mh.swap(smallest, index)
		mh.Heapify(smallest, n)
	}
}

func (mh *CustomMinHeap) top() *Cell {
	return (*mh)[0]
}

func (mh *CustomMinHeap) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}
