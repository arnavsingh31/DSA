package heap

/*
	GFG
	TC--->O(n^2logk)
	SC--->O(k)
*/
func KthLargestSum(arr []int, k int) int {
	minHeap := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]

			if len(minHeap) < k {
				InsertInMinHeap(&minHeap, sum)
			} else {
				if minHeap[0] < sum {
					DeleteFromMinHeap(&minHeap)
					InsertInMinHeap(&minHeap, sum)
				}
			}

		}
	}

	return minHeap[0]
}
