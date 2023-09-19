package heap

// Insert Element in max-heap
func InsertInMaxHeap(arr *[]int, val int) {
	// first insert at end
	*arr = append(*arr, val)
	index := len(*arr) - 1

	//then find correct position of the val {here using max heap to find val position}
	for index > 0 {
		parent := (index - 1) / 2

		if (*arr)[parent] < (*arr)[index] {
			(*arr)[parent], (*arr)[index] = (*arr)[index], (*arr)[parent]
			index = parent
		} else {
			return
		}
	}
}

// Insert Element in min-heap
func InsertInMinHeap(arr *[]int, val int) {
	*arr = append(*arr, val)
	index := len(*arr) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*arr)[parent] > (*arr)[index] {
			(*arr)[parent], (*arr)[index] = (*arr)[index], (*arr)[parent]
			index = parent
		} else {
			return
		}
	}
}

// we can delete like this also.
func deleteFromHeap(arr *[]int) {

	// first copy last node to first node
	(*arr)[1] = (*arr)[len(*arr)-1]

	// delete last node
	*arr = (*arr)[:len(*arr)-1]
	n := len(*arr)
	// now find the correct position of root in heap
	index := 1
	for index < n {
		left := 2 * index
		right := 2*index + 1
		swapIndex := index

		if left < n && (*arr)[swapIndex] < (*arr)[left] {
			swapIndex = left
		}

		if right < n && (*arr)[swapIndex] < (*arr)[right] {
			swapIndex = right
		}

		if swapIndex != index {
			(*arr)[index], (*arr)[swapIndex] = (*arr)[swapIndex], (*arr)[index]
		} else {
			return
		}

	}
}

// Delete top(root) element from max-heap
func DeleteFromMaxHeap(arr *[]int) {
	lastIndex := len(*arr) - 1

	// swap first and last
	(*arr)[0], (*arr)[lastIndex] = (*arr)[lastIndex], (*arr)[0]
	*arr = (*arr)[:lastIndex]

	//max-heapify
	MaxHeapify(arr, 0, len(*arr)-1)

}

// Delete top(root) element from min-heap
func DeleteFromMinHeap(arr *[]int) {
	lastIndex := len(*arr) - 1

	// swap first and last
	(*arr)[0], (*arr)[lastIndex] = (*arr)[lastIndex], (*arr)[0]
	*arr = (*arr)[:lastIndex]

	//min-heapify
	MinHeapify(arr, 0, len(*arr)-1)

}

// Max-Heapify {heapify algorithm}
func MaxHeapify(arr *[]int, i, n int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left <= n && (*arr)[largest] < (*arr)[left] {
		largest = left
	}

	if right <= n && (*arr)[largest] < (*arr)[right] {
		largest = right
	}

	if largest != i {
		(*arr)[i], (*arr)[largest] = (*arr)[largest], (*arr)[i]
		MaxHeapify(arr, largest, n)
	}
}

// Min-Heapify {same as max-heapify, just reverse in-equalities}
func MinHeapify(arr *[]int, i, n int) {
	left := 2*i + 1
	right := 2*i + 2
	smallest := i

	if left <= n && (*arr)[smallest] > (*arr)[left] {
		smallest = left
	}

	if right <= n && (*arr)[smallest] > (*arr)[right] {
		smallest = right
	}

	if smallest != i {
		(*arr)[i], (*arr)[smallest] = (*arr)[smallest], (*arr)[i]
		MinHeapify(arr, smallest, n)
	}
}

/*
	Heap Sort
	TC-->O(nlogn)
	1. swap first with last
	2. now apply heapify algo to move root val to its correct position in array.
*/
func HeapSort(arr *[]int, n int) {

	for i := n; i > 0; i-- {
		// swap first and last element, decrement the size
		(*arr)[1], (*arr)[i] = (*arr)[i], (*arr)[1]
		n--

		MaxHeapify(arr, n, 1)
	}
}

// heapify algorithm TC--->O(n) how?
// arr := []int{-1, 54, 53, 55, 52, 50}
// n := len(arr)
// for i := (n - 1) / 2; i > 0; i-- {
// 	heapify2(&arr, i)
// }

// 	arr := []int{-1, 60, 55, 70, 45, 50}
// 	HeapSort(&arr, len(arr)-1)
// 	log.Print(arr)
// }

/*
----------------------------------------------MIN-HEAP IMPLEMENTATION----------------------------------------
*/
type MinHeap []int

// Inserting an element in heap.
func (mh *MinHeap) push(val int) {
	*mh = append(*mh, val)

	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent] > (*mh)[index] {
			mh.swap(parent, index)
			index = parent
		} else {
			return
		}
	}
}

// Deleting top from heap.
func (mh *MinHeap) pop() {
	lastIndex := len(*mh) - 1

	//swap first and last nodes
	mh.swap(0, lastIndex)

	// pop last node
	*mh = (*mh)[:lastIndex]

	// heapify root index i.e 0
	mh.heapify(0, len(*mh)-1)
}

// Returns top of the heap
func (mh *MinHeap) top() int {
	return (*mh)[0]
}

func (mh *MinHeap) heapify(index int, n int) {
	left := 2*index + 1
	right := 2*index + 2
	smallest := index

	if left <= n && (*mh)[smallest] > (*mh)[left] {
		smallest = left
	}

	if right <= n && (*mh)[smallest] > (*mh)[right] {
		smallest = right
	}

	if smallest != index {
		mh.swap(smallest, index)
		mh.heapify(smallest, n)
	}
}

func (mh *MinHeap) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

/*
----------------------------------------------MAX-HEAP IMPLEMENTATION----------------------------------------
*/
type MaxHeap []int

// Inserting an element in heap.
func (mh *MaxHeap) push(val int) {
	*mh = append(*mh, val)

	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent] < (*mh)[index] {
			mh.swap(parent, index)
			index = parent
		} else {
			return
		}
	}
}

// Deleting top from heap.
func (mh *MaxHeap) pop() {
	lastIndex := len(*mh) - 1

	//swap first and last nodes
	mh.swap(0, lastIndex)

	// pop last node
	*mh = (*mh)[:lastIndex]

	// heapify root index i.e 0
	mh.heapify(0, len(*mh)-1)
}

// Returns top of the heap
func (mh *MaxHeap) top() int {
	return (*mh)[0]
}

func (mh *MaxHeap) heapify(index int, n int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left <= n && (*mh)[largest] < (*mh)[left] {
		largest = left
	}

	if right <= n && (*mh)[largest] < (*mh)[right] {
		largest = right
	}

	if largest != index {
		mh.swap(largest, index)
		mh.heapify(largest, n)
	}
}

func (mh *MaxHeap) swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}
