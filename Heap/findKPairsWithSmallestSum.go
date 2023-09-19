package heap

/*
	LC #373
	TC--->O(min{klogk, m*nlogmn})
	SC--->O(min{k, m*n})
*/

type HeapNode struct {
	sum int
	i   int
	j   int
}

type MinSumHeap []*HeapNode

type Pair struct {
	i, j int
}

func FindKPairs(arr1, arr2 []int, k int) [][]int {
	minHeap := make(MinSumHeap, 0)
	ans := [][]int{}
	visited := make(map[Pair]bool)

	minHeap.push(&HeapNode{sum: arr1[0] + arr2[0], i: 0, j: 0})
	visited[Pair{0, 0}] = true

	for len(minHeap) > 0 && k > 0 {
		top := minHeap.top()
		minHeap.pop()

		i := top.i
		j := top.j

		ans = append(ans, []int{arr1[i], arr2[j]})

		if i+1 < len(arr1) && !visited[Pair{i + 1, j}] {
			minHeap.push(&HeapNode{sum: arr1[i+1] + arr2[j], i: i + 1, j: j})
			visited[Pair{i + 1, j}] = true
		}

		if j+1 < len(arr2) && !visited[Pair{i, j + 1}] {
			minHeap.push(&HeapNode{sum: arr1[i] + arr2[j+1], i: i, j: j + 1})
			visited[Pair{i, j + 1}] = true
		}
		k--
	}

	return ans
}

func (mh *MinSumHeap) push(hn *HeapNode) {
	*mh = append(*mh, hn)

	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent].sum > (*mh)[index].sum {
			mh.swap(parent, index)
			index = parent
		} else {
			return
		}
	}
}

// Deleting top from heap.
func (mh *MinSumHeap) pop() {
	lastIndex := len(*mh) - 1

	//swap first and last nodes
	mh.swap(0, lastIndex)

	// pop last node
	*mh = (*mh)[:lastIndex]

	// heapify root index i.e 0
	mh.heapify(0, len(*mh)-1)
}

// Returns top of the heap
func (mh *MinSumHeap) top() *HeapNode {
	return (*mh)[0]
}

func (mh *MinSumHeap) heapify(index int, n int) {
	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left <= n && (*mh)[smallest].sum > (*mh)[left].sum {
		smallest = left
	}

	if right <= n && (*mh)[smallest].sum > (*mh)[right].sum {
		smallest = right
	}

	if smallest != index {
		mh.swap(smallest, index)
		mh.heapify(smallest, n)
	}
}

func (mh *MinSumHeap) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}
