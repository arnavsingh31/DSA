package matrices

/*
	LC #766
	TC--->O(m*n)
	SC--->O(1)
*/
func IsToeplitzMatrix(matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}
