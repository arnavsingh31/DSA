package main

/*
	LC #867
	TC--->O(m*n), m= rows, n= cols
	SC--->O(m*n)
*/
func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	res := make([][]int, cols)

	for i := range res {
		res[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
