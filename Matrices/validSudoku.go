package main

/*
	LC #36
	TC--->O(9^2)
	SC--->O(9^2)
*/
func isValidSudoku(board [][]byte) bool {
	rows := len(board)
	cols := len(board[0])

	rowMap := make(map[int]map[byte]struct{}, rows-1)
	colMap := make(map[int]map[byte]struct{}, cols-1)
	gridMap := make(map[Pos]map[byte]struct{}, 0)

	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[byte]struct{}, 0)
		colMap[i] = make(map[byte]struct{}, 0)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			gridMap[Pos{i / 3, j / 3}] = make(map[byte]struct{}, 0)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '.' {
				continue
			}

			// exist in 3x3 grid or not
			if _, exist := gridMap[Pos{i / 3, j / 3}][board[i][j]]; exist {
				return false
			}

			// exist in a row or not
			if _, exist := rowMap[i][board[i][j]]; exist {
				return false
			}

			// exist in col or not
			if _, exist := colMap[j][board[i][j]]; exist {
				return false
			}

			gridMap[Pos{i / 3, j / 3}][board[i][j]] = struct{}{}
			rowMap[i][board[i][j]] = struct{}{}
			colMap[j][board[i][j]] = struct{}{}
		}
	}
	return true
}
