package graph

/*
	LC #733
	TC--->O(n+2E)
	SC--->O(n)
*/
func FloodFill(image [][]int, sr, sc, color int) [][]int {
	rows := len(image)
	cols := len(image[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := make([][]int, rows)

	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ans[i][j] = image[i][j]
		}
	}

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	bfs := func(i, j int) {
		ans[i][j] = color
		queue := []Cell{{i, j}}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			x, y := cell.row, cell.col

			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]

				if isCellValid(newX, newY) && ans[newX][newY] != color && image[x][y] == image[newX][newY] {
					queue = append(queue, Cell{newX, newY})
					ans[newX][newY] = color
				}
			}
		}
	}

	bfs(sr, sc)
	return ans
}

type Cell struct {
	row, col int
}
