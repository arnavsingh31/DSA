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
	directions := [][]int{{1, 0}, {0, -1}, {0, 1}, {-1, 0}}
	move := []string{"D", "L", "R", "U"}

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

		for k, dir := range directions {
			newX, newY := i+dir[0], j+dir[1]

			if isCellValid(newX, newY) && !visited[Pos{newX, newY}] {
				dfs(newX, newY, path+move[k])
			}
		}

		visited[Pos{i, j}] = false

	}

	dfs(0, 0, "")

	return ans
}
