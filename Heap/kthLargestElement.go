package heap

/*
LC #217
TC--->O()
SC--->O()
*/
func KthLargest(arr []int, k int) int {
	// heapify the given array
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		MaxHeapify(&arr, i, len(arr)-1)
	}

	// remove root k-1 times and also heapify after each deletion.
	for i := 0; i < k-1; i++ {
		DeleteFromMaxHeap(&arr)
	}

	// kth largest element will be our root element
	return arr[0]
}

/*
	Algo:
	1. first build a min-heap of first k element.
	2. then for k-1 elements, if element is > heap-top then delete top and insert this element.
	3. after processiong all k-1 elements, our kth largest element will be at top of our heap.
*/
func KthLargest2(arr []int, k int) int {
	minHeap := make([]int, 0)

	for i := 0; i < k; i++ {
		InsertInMinHeap(&minHeap, arr[i])
	}

	for i := k; i < len(arr); i++ {
		if arr[i] > minHeap[0] {
			DeleteFromMinHeap(&minHeap)
			InsertInMinHeap(&minHeap, arr[i])
		}
	}

	return minHeap[0]
}
