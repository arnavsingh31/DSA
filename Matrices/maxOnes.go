package main

/*
	LC #2643
	TC--->O(m*n)
	SC--->O(1)
*/
func maxOnesRow(matrix [][]int) []int {
	ans := make([]int, 2)

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		max := 0
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				max++
			}
		}

		if max > ans[1] {
			ans[0] = i
			ans[1] = max
		}
	}

	return ans
}
