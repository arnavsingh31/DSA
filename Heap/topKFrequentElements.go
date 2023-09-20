package heap

/*
	LC #347
	TC--->O(DlogD{for insertion in heap} + kLogD{for deletion in heap}) {D-> number of distinct elements}
	SC--->O(D) {map + heap}
*/
func FindTopKFrequentElements(arr []int, k int) []int {
	if len(arr) == k || len(arr) == 1 {
		return arr
	}

	numToCountMap := make(map[int]int)
	maxHeap := make(ElementMaxHeap, 0)
	ans := []int{}

	// populate map
	for i := 0; i < len(arr); i++ {
		numToCountMap[arr[i]] += 1
	}

	// populate maxHeap from map
	for num, freq := range numToCountMap {
		maxHeap.push(&Element{
			data:  num,
			count: freq,
		})
	}

	for k > 0 {
		ans = append(ans, maxHeap.top().data)
		maxHeap.pop()
		k--
	}

	return ans
}

type Element struct {
	data  int
	count int
}

type ElementMaxHeap []*Element

func (mh *ElementMaxHeap) push(e *Element) {
	*mh = append(*mh, e)

	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent].count < (*mh)[index].count {
			mh.swap(parent, index)
			index = parent
		} else {
			return
		}
	}
}

// Deleting top from heap.
func (mh *ElementMaxHeap) pop() {
	lastIndex := len(*mh) - 1

	//swap first and last nodes
	mh.swap(0, lastIndex)

	// pop last node
	*mh = (*mh)[:lastIndex]

	// heapify root index i.e 0
	mh.heapify(0, len(*mh)-1)
}

// Returns top of the heap
func (mh *ElementMaxHeap) top() *Element {
	return (*mh)[0]
}

func (mh *ElementMaxHeap) heapify(index int, n int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left <= n && (*mh)[largest].count < (*mh)[left].count {
		largest = left
	}

	if right <= n && (*mh)[largest].count < (*mh)[right].count {
		largest = right
	}

	if largest != index {
		mh.swap(largest, index)
		mh.heapify(largest, n)
	}
}

func (mh *ElementMaxHeap) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

/*
	Using Bucket Sort
	TC--->O(n)
	SC--->O(D + n ) {D-> number of distinct element in map and n for bucket}
*/
func FindTopKFrequentElements2(arr []int, k int) []int {
	numToCountMap := make(map[int]int)
	bucket := make([][]int, len(arr)+1)
	ans := []int{}

	// populate map
	for i := 0; i < len(arr); i++ {
		numToCountMap[arr[i]] += 1
	}

	for key, val := range numToCountMap {
		bucket[val] = append(bucket[val], key)
	}

	for i := len(bucket) - 1; i > 0 && k > 0; i-- {
		if len(bucket[i]) > 0 {
			for j := 0; j < len(bucket[i]); j++ {
				ans = append(ans, bucket[i][j])
				k--
			}
		}
	}

	return ans

}
