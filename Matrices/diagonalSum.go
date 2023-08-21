package main

/*
	LC #1572
	TC--->O(m+n), we are only visiting elements on diagonals.
	SC--->O(1)
*/
func diagonalSum(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	total := rows * cols

	startingRow := 0
	endingRow := rows - 1
	startingCol := 0
	endingCol := cols - 1

	sum := 0

	for startingRow <= endingRow && startingCol <= endingCol {
		sum += matrix[startingRow][startingCol] + matrix[startingRow][endingCol] + matrix[endingRow][startingCol] + matrix[endingRow][endingCol]

		startingRow++
		startingCol++
		endingRow--
		endingCol--
	}

	if total%2 == 0 {
		return sum
	}

	return sum - 3*matrix[startingRow][startingCol]

}

// using only 2 pointers.
func diagonalSum2(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	startingCol := 0
	endingCol := cols - 1

	sum := 0
	for row := 0; row < rows && startingCol <= cols-1 && endingCol >= 0; row++ {
		if startingCol == endingCol {
			sum += matrix[row][startingCol]
		} else {
			sum += matrix[row][startingCol] + matrix[row][endingCol]
		}

		startingCol++
		endingCol--
	}

	return sum
}
