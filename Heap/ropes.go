package heap

/*
	GFG-problem
	TC--->O(nlogn)
	SC--->O(n) {if we use priority queue}, since i used input array for priority-queue/min-heap, so O(1).
*/
func MinimumCostofRopes(arr []int) int {
	size := len(arr) - 1
	// build a min-heap
	for i := (size - 1) / 2; i >= 0; i-- {
		MinHeapify(&arr, i, size)
	}

	cost := 0
	for len(arr) > 1 {
		min1 := arr[0]
		DeleteFromMinHeap(&arr)
		min2 := arr[0]
		DeleteFromMinHeap(&arr)

		sum := min1 + min2
		cost += sum
		InsertInMinHeap(&arr, sum)
	}

	return cost
}

/*

		2
	   / \
	  3	  4
	 /
	6

	cost1 = 2 + 3 ->5
	cost2 = 5 + 4 ->9
	cost3 = 6 + 9 -> 15

	total cost = cost1 + cost2 + cost3 = 29
*/
