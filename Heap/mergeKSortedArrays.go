package heap

/*
	GFG
	TC--->O(klogk + n*klogk)~ O(n*klogk)
	SC--->O(k + n*k) {min-heap + ans array} ~ O(n*k)
*/

type Cell struct {
	row int
	col int
	val int
}

func MergeKSortedArray(matrix [][]int) []int {
	rows := len(matrix)
	ans := make([]int, 0)

	minHeap := make(MinHeap, 0)

	// O(klogk)
	for i := 0; i < rows; i++ {
		// insert first element of all k array in minHeap
		temp := &Cell{
			row: i,
			col: 0,
			val: matrix[i][0],
		}
		minHeap.push(temp)
	}

	// elements left to be processed = n*k - k, where n*k are total elements in matrix and we have already
	// added k elements in minheap, so this ~n*k and insertion and deletion will take logK so overall
	// complexity of loop below is O(n*klogk).
	for len(minHeap) > 0 {
		temp := minHeap.top()
		ans = append(ans, temp.val)
		minHeap.pop()

		i := temp.row
		j := temp.col

		if j+1 < len(matrix[i]) {
			nextCell := &Cell{
				row: i,
				col: j + 1,
				val: matrix[i][j+1],
			}

			minHeap.push(nextCell)
		}
	}
	return ans
}
