package heap

import "container/heap"

/*
	LC #378 {Note: rows = col = n, as matrix is square}
	TC--->O(n*n + nlogk)
	SC--->O(n*n + k) {array + heap}
*/
func KthSmallestInMatrix(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	arr := make([]int, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			arr = append(arr, matrix[i][j])
		}
	}

	return KthSmallest(arr, k)
}

/*
	Kth smallest element can also be seen as largest element in a priority queue of size k
	By this logic we pop elements from priority queue when its size becomes greater than k
	thus top of priority queue is kth smallest element in matrix
	TC--->O((n*n)logk)
	SC--->O(k)
*/
func KthSmallestInMatrix2(matrix [][]int, k int) int {
	maxHeap := make(MaxHeap, 0)
	rows := len(matrix)
	cols := len(matrix[0])
	var i, j int

	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			maxHeap.push(matrix[i][j])
			if len(maxHeap) > k {
				maxHeap.pop()
			}
		}
	}

	return maxHeap.top()
}

/*
	TC--->O(nlogn)
	SC--->O(n)
	Using Min-Heap
*/
func KthSmallestInMatrix3(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	minHeap := make(CustomMinHeap, 0)

	// O(nlogn)
	for i := 0; i < rows; i++ {
		minHeap.push(&Cell{
			row: i,
			col: 0,
			val: matrix[i][0],
		})
	}

	// O(klogn)
	for i := 0; i < k-1; i++ {
		top := minHeap.top()
		minHeap.pop()

		if top.col+1 < cols {
			minHeap.push(&Cell{
				row: top.row,
				col: top.col + 1,
				val: matrix[top.row][top.col+1],
			})
		}
	}

	return minHeap.top().val
}

/*
	priority_queue using heap library, we need to implement basic interface. after that we can use heap.
*/
func KthSmallestInMatrix4(matrix [][]int, k int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	pq := make(PriorityQueue, rows)

	// O(n)
	for i := 0; i < rows; i++ {
		pq[i] = &Item{
			x:    i,
			y:    0,
			data: matrix[i][0],
		}
	}

	// O(n)
	heap.Init(&pq)

	// O(klogn)
	for i := 0; i < k-1; i++ {
		top := heap.Pop(&pq).(*Item)

		if top.y+1 < cols {
			heap.Push(&pq, &Item{
				x:    top.x,
				y:    top.y + 1,
				data: matrix[top.x][top.y+1],
			})
		}
	}

	return heap.Pop(&pq).(*Item).data
}
