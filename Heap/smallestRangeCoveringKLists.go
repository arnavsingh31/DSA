package heap

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #632
TC--->O(klogk+nklogk)~ O(n*klogk)
SC--->O(k)
Using min-heap
*/
func SmallestRangeCoveringKLists(matrix [][]int) []int {
	smallestRange := make([]int, 2)

	minHeap := make(MinHeap, 0)
	max := math.MinInt
	min := math.MaxInt

	for i := 0; i < len(matrix); i++ {
		max = util.Max(max, matrix[i][0])
		min = util.Min(min, matrix[i][0])
		temp := &Cell{
			row: i,
			col: 0,
			val: matrix[i][0],
		}
		minHeap.push(temp)
	}

	start, end := min, max

	for len(minHeap) > 0 {
		top := minHeap.top()
		min = top.val
		minHeap.pop()

		if max-min < end-start {
			start = min
			end = max
		}

		if top.col+1 < len(matrix[top.row]) {
			next := &Cell{
				row: top.row,
				col: top.col + 1,
				val: matrix[top.row][top.col+1],
			}
			max = util.Max(max, matrix[top.row][top.col+1])
			minHeap.push(next)
		} else {
			break
		}
	}

	smallestRange[0] = start
	smallestRange[1] = end

	return smallestRange
}
