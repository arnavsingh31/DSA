package matrices

/*
	LC #1351
	TC--->O(m+n)
	SC--->O(1)
*/
func CountNegative(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	startingRow := 0
	endingCol := cols - 1

	count := 0

	for startingRow < rows && endingCol >= 0 {
		num := grid[startingRow][endingCol]

		if num < 0 {
			count += rows - startingRow
			endingCol--
		}

		if num >= 0 {
			startingRow++
		}
	}
	return count
}
