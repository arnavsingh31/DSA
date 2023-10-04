package matrices

/*
	LC #59
	TC--->O(m*n)
	SC--->O(1)
*/
func GenerateMatrix(n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	startingRow, startingCol := 0, 0
	endingRow, endingCol := n-1, n-1

	counter := 1

	for counter <= n*n {
		for i := startingCol; i <= endingCol && counter <= n*n; i++ {
			matrix[startingRow][i] = counter
			counter++
		}
		startingRow++

		for i := startingRow; i <= endingRow && counter <= n*n; i++ {
			matrix[i][endingCol] = counter
			counter++
		}
		endingCol--

		for i := endingCol; i >= startingCol && counter <= n*n; i-- {
			matrix[endingRow][i] = counter
			counter++
		}
		endingRow--

		for i := endingRow; i >= startingRow && counter <= n*n; i-- {
			matrix[i][startingCol] = counter
			counter++
		}
		startingCol++
	}

	return matrix
}
