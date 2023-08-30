package main

/*
	LC #79
	TC--->O(m*n*4^len(word))
	SC--->O(m*n)
*/
func wordSearch(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	path := make(map[Pos]struct{}, 0)
	charArr := []byte(word)

	var dfs func(int, int, int) bool
	dfs = func(i, j, currWordIndex int) bool {
		if currWordIndex == len(word) {
			return true
		}

		if _, exist := path[Pos{i, j}]; exist {
			return false
		}
		// if our index is out of bound or char at currIndex is not equal to board[i][j] return false
		if i < 0 || j < 0 || i >= rows || j >= cols || board[i][j] != charArr[currWordIndex] {
			return false
		}

		path[Pos{i, j}] = struct{}{}
		res := dfs(i+1, j, currWordIndex+1) || dfs(i, j+1, currWordIndex+1) || dfs(i-1, j, currWordIndex+1) || dfs(i, j-1, currWordIndex+1)
		delete(path, Pos{i, j})
		return res
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == charArr[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
