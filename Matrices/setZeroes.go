package matrices

/*
	LC #73
	TC--->O(m*n)
	SC--->O(m+n), we use separate row and col array to mark corresponding index of both to 1 (means we have
	to mark entire row/col of matrix to 0). And later we use these array to set matrix elements to 0.
*/
func SetZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	rowArr := make([]int, rows)
	colArr := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				rowArr[i] = 1
				colArr[j] = 1
			}
		}
	}

	for i := 0; i < rows; i++ {
		if rowArr[i] == 1 {
			for j := 0; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < cols; j++ {
		if colArr[j] == 1 {
			for i := 0; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

/*
	TC--->O(m*n)
	SC--->O(1)
	In this approach, we can just improve the space complexity. So, instead of using two extra array row
	and col, we will use the 1st row and 1st column of the given matrix to keep a track of the cells that need
	to be marked with 0. But here comes a problem. If we try to use the 1st row and 1st column to serve the
	purpose, the cell matrix[0][0] is taken twice. To solve this problem we will take an extra variable col0
	(we can use row0 also, as implemented below) initialized with 1. Now the entire 1st column of the matrix will
	serve the purpose of the row array. And the 1st row from (0,1) to (0,m-1) with the col0 variable will
	serve the purpose of the col array. In my case, entire first row from (0,0) to (0, len(matrix[0])-1) will be our
	col array and first column from (1,0) to (len(matrix)-1, 0) and row0 will be our row array.
*/
func SetZeroes2(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	row0 := 1

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				if i == 0 {
					row0 = 0
				} else {
					matrix[i][0] = 0
				}
			}
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// you need to check matrix[0][0] == 0 BEFORE* row0 == 0 because in case row0 is 0 but matrix[0][0] != 0 then
	// we will overide the (0,0) cell in matrix and then we will also make enitre first col = 0.
	if matrix[0][0] == 0 {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}

	if row0 == 0 {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}

}
