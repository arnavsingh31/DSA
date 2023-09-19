package backtracking

/*
	TC--->O(4^(n^2))
	SC--->O(n*m)
*/
type Pos struct {
	row int
	col int
}

func FindPathsInMaze(grid [][]int) []string {
	rows := len(grid)
	cols := len(grid[0])
	ans := []string{}
	visited := make(map[Pos]bool)

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols && grid[i][j] == 1
	}

	var dfs func(int, int, string)
	dfs = func(i, j int, path string) {
		if i == rows-1 && j == cols-1 {
			ans = append(ans, path)
			return
		}

		visited[Pos{i, j}] = true

		if isCellValid(i+1, j) && !visited[Pos{i + 1, j}] {
			dfs(i+1, j, path+"D")
		}

		if isCellValid(i, j-1) && !visited[Pos{i, j - 1}] {
			dfs(i, j-1, path+"L")
		}

		if isCellValid(i, j+1) && !visited[Pos{i, j + 1}] {
			dfs(i, j+1, path+"R")
		}

		if isCellValid(i-1, j) && !visited[Pos{i - 1, j}] {
			dfs(i-1, j, path+"U")
		}

		visited[Pos{i, j}] = false

	}

	dfs(0, 0, "")

	return ans
}
