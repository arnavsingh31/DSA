package backtracking

/*
	LC #37
	TC--->O(9^(n*n))
	SC--->O(1)
*/
func SolveSudoku(board [][]byte) {
	solve(&board)
}

func solve(board *[][]byte) bool {
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			if (*board)[i][j] == '.' {
				for k := byte('1'); k <= '9'; k++ {
					if validNumForCell(board, i, j, k) {
						(*board)[i][j] = k
						if solve(board) {
							return true
						}
						// backtrack
						(*board)[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func validNumForCell(board *[][]byte, row, col int, char byte) bool {
	for i := 0; i < 9; i++ {
		if (*board)[row][i] == char {
			return false
		}

		if (*board)[i][col] == char {
			return false
		}
	}

	rowStart, rowEnd := (row/3)*3, (row/3+1)*3
	colStart, colEnd := (col/3)*3, (col/3+1)*3
	for i := rowStart; i < rowEnd; i++ {
		for j := colStart; j < colEnd; j++ {
			if (*board)[i][j] == char {
				return false
			}
		}
	}

	return true
}
