package matrices

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #85
	TC--->O(m*(n+n)) {m rows times {n for second for loop + n for TC of largestHistogramArea}}
	SC--->O(m*n) {for rach row {m} we use stack space{n}}
*/
func MaximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		maxArea = util.Max(maxArea, largestHistogramArea(heights))
	}

	return maxArea
}

func largestHistogramArea(heights []int) int {
	prevSmallerHeight := make([]int, len(heights))
	nextSmallerHeight := make([]int, len(heights))

	for i := 0; i < len(heights); i++ {
		prevSmallerHeight[i] = -1
		nextSmallerHeight[i] = len(heights)
	}

	stack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextSmallerHeight[topIndex] = i
		}

		if len(stack) > 0 {
			prevSmallerHeight[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		width := nextSmallerHeight[i] - (prevSmallerHeight[i] + 1)
		maxArea = util.Max(maxArea, heights[i]*width)
	}

	return maxArea
}
