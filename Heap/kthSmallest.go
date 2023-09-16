package heap

/*
	GFG {Use max-heap to find kth smallest}
*/
func KthSmallest(arr []int, k int) int {
	heap := make([]int, 0)

	for i := 0; i < k; i++ {
		InsertInMaxHeap(&heap, arr[i])
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			DeleteFromMaxHeap(&heap)
			InsertInMaxHeap(&heap, arr[i])
		}
	}

	return heap[0]
}
