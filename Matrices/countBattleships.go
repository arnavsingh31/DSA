package main

/*
	LC #419
	TC--->O(m*n)
	SC--->O(m+n)
*/
func countBattleships(board [][]byte) int {
	rows := len(board)
	cols := len(board[0])
	count := 0

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i >= rows || j >= cols || board[i][j] == '.' {
			return
		}

		board[i][j] = '1'
		dfs(i+1, j)
		dfs(i, j+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'X' {
				count++
				dfs(i, j)
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '1' {
				board[i][j] = 'X'
			}
		}
	}

	return count
}
