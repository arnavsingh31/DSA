package main

/*
	LC #48
	TC--->O(m*n)
	SC--->O(1)
*/
func rotate(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	// find transpose of matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// now change cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols/2; j++ {
			matrix[i][j], matrix[i][cols-1-j] = matrix[i][cols-1-j], matrix[i][j]
		}
	}

}
