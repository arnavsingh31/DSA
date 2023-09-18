package heap

/*
	LC #23
	TC--->O(Nlogk) {N = total number of nodes in linked list, k is size of list}
	SC--->O(K) {min-heap}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKSortedList(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var head, tail *ListNode
	minHeap := make(ListMinHeap, 0)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			minHeap.push(lists[i])
		}
	}

	for len(minHeap) > 0 {
		top := minHeap.top()
		minHeap.pop()

		if top.Next != nil {
			minHeap.push(top.Next)
		}

		if head == nil {
			head = top
			tail = top
		} else {
			tail.Next = top
			tail = top
		}
	}

	return head
}

type ListMinHeap []*ListNode

func (mh *ListMinHeap) push(node *ListNode) {
	*mh = append(*mh, node)

	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent].Val > (*mh)[index].Val {
			mh.swap(parent, index)
			index = parent
		} else {
			return
		}
	}
}

func (mh *ListMinHeap) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *ListMinHeap) top() *ListNode {
	return (*mh)[0]
}

func (mh *ListMinHeap) pop() {
	lastIndex := len(*mh) - 1

	// swap first and last
	mh.swap(0, lastIndex)
	*mh = (*mh)[:lastIndex]

	mh.heapify(0, len(*mh)-1)
}

func (mh *ListMinHeap) heapify(index, n int) {
	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left <= n && (*mh)[smallest].Val > (*mh)[left].Val {
		smallest = left
	}

	if right <= n && (*mh)[smallest].Val > (*mh)[right].Val {
		smallest = right
	}

	if smallest != index {
		mh.swap(smallest, index)
		mh.heapify(smallest, n)
	}
}
