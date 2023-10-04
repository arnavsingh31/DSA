package matrices

/*
	LC #54 (Simulate the spiral pattern)
	TC--->O(n*m)
	SC--->O(1), without considering space of ans
*/
func SpiralMatrix(arr [][]int) []int {
	rows := len(arr)
	cols := len(arr[0])
	ans := []int{}

	total := rows * cols
	count := 0

	starting_row := 0
	starting_col := 0
	ending_row := rows - 1
	ending_col := cols - 1

	for count < total {
		// traversing first row
		for i := starting_col; i <= ending_col; i++ {
			ans = append(ans, arr[starting_row][i])
			count++
		}
		starting_row++

		// traversing last col
		for i := starting_row; i <= ending_row; i++ {
			ans = append(ans, arr[i][ending_col])
			count++
		}
		ending_col--

		// tarversing last row
		for i := ending_col; i >= starting_col; i-- {
			ans = append(ans, arr[ending_row][i])
			count++
		}
		ending_row--

		// traversing first col
		for i := ending_row; i >= starting_row; i-- {
			ans = append(ans, arr[i][starting_col])
			count++
		}
		starting_col++
	}

	return ans
}
