package main

/*
	LC #419
	TC--->O(m*n)
	SC--->O(m+n), call stack space
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

/*
	TC--->O(m*n)
	SC--->O(1)
	For each cell, check if it contains a battleship ('X').

	- If the cell contains a battleship ('X'), we need to determine if it's a starting cell of a battleship.
	To do this, we check two conditions:
	Condition 1: Check if the cell above is empty ('.') or if it's the first row.
	Condition 2: Check if the cell to the left is empty ('.') or if it's the first column.

	If both conditions are met, increment the cnt counter. This means we found a new battleship.

*/
func countBattleships2(board [][]int) int {
	rows := len(board)
	cols := len(board[0])
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'X' {
				if (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
					count++
				}
			}
		}
	}
	return count
}
