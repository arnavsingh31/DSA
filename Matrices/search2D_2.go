package matrices

/*
	LC #240
	TC--->O(log(m*n))
	SC--->O(1)

	Here, we need to begin at end of a row. and then apply binary search algorithm.
*/
func Search2D2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	rowIndex := 0
	colIndex := cols - 1

	for rowIndex < rows && colIndex >= 0 {
		element := matrix[rowIndex][colIndex]

		if element == target {
			return true
		}

		if element > target {
			colIndex--
		}

		if element < target {
			rowIndex++
		}
	}
	return false
}
