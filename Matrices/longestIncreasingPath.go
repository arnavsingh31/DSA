package matrices

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #329
	TC--->O(m*n)
	SC--->O(m*n)
*/
func LongestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	maxPath := 0

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	mem := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mem[i] = make([]int, cols)
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if mem[i][j] > 0 {
			return mem[i][j]
		}

		path := 0
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]

			if isCellValid(x, y) && matrix[x][y] > matrix[i][j] {
				path = util.Max(path, dfs(x, y))
			}
		}

		mem[i][j] = path + 1
		return mem[i][j]
	}

	/* DRIVER CODE */
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			temp := dfs(i, j)
			maxPath = util.Max(maxPath, temp)
		}
	}

	return maxPath
}
