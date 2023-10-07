package matrices

/*
	LC #130
	TC--->O(m*n)
	SC--->O(m*n)
*/
func Solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	dfs := func(i, j int) {
		stack := []Cell{{i, j}}

		for len(stack) > 0 {
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			x, y := cell.row, cell.col
			board[x][y] = 'F'

			if x > 0 && board[x-1][y] == 'O' {
				stack = append(stack, Cell{x - 1, y})
			}
			if y > 0 && board[x][y-1] == 'O' {
				stack = append(stack, Cell{x, y - 1})
			}
			if x < rows-1 && board[x+1][y] == 'O' {
				stack = append(stack, Cell{x + 1, y})
			}
			if y < cols-1 && board[x][y+1] == 'O' {
				stack = append(stack, Cell{x, y + 1})
			}
		}
	}

	// traversing around border cell and if cell is 'O' then apply dfs to that cell and change that cell val to
	// 'F'.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' && (i == 0 || i == rows-1 || j == 0 || j == cols-1) {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}
}
