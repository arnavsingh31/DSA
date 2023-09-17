package heap

/*
	Merge 2 binary max heap
	TC--->O(n)
*/
func MergeMaxHeap(heap1, heap2 []int) []int {
	// merge given 2 arrays
	heap1 = append(heap1, heap2...)
	size := len(heap1) - 1

	// now heapify the combined array
	for i := (size - 1) / 2; i >= 0; i-- {
		MaxHeapify(&heap1, i, size)
	}

	return heap1
}
