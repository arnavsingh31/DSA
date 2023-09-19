package backtracking

/*
	LC #51
	TC--->O(n!)
	SC--->O(n*n)
*/
func PlaceNQueens(n int) [][]string {
	board := make([][]string, n)
	ans := make([][]string, 0)

	// fill board with empty spaces
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}

	solveNQ(&board, &ans, 0, n)
	return ans
}

func solveNQ(board, ans *[][]string, col, n int) {
	if col == n {
		//copy board -> ans
		addAns(board, ans, n)
		return
	}

	for row := 0; row < n; row++ {
		if isQueenSafe(row, col, n, board) {
			(*board)[row][col] = "Q"
			solveNQ(board, ans, col+1, n)
			(*board)[row][col] = "."
		}
	}
}

func addAns(board, ans *[][]string, n int) {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		temp := ""
		for j := 0; j < n; j++ {
			temp += (*board)[i][j]
		}
		arr[i] += temp
	}

	*ans = append(*ans, arr)
}

func isQueenSafe(i, j, n int, board *[][]string) bool {
	x, y := i, j
	//checking previous cols along same row
	for y >= 0 {
		if (*board)[x][y] == "Q" {
			return false
		}
		y--
	}

	x, y = i, j
	// checking upper left diagonal cells
	for x >= 0 && y >= 0 {
		if (*board)[x][y] == "Q" {
			return false
		}
		x--
		y--
	}

	x, y = i, j
	// checking bottom left diagonal cells
	for x < n && y >= 0 {
		if (*board)[x][y] == "Q" {
			return false
		}
		x++
		y--
	}

	return true
}
