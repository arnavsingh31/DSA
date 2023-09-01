package main

/*
	LC #695
	TC--->O(m*n)
	SC--->O(m*n)
*/
func maxAreaIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	maxArea := 0

	bfs := func(i, j int) {
		queue := []Pos{{i, j}}
		islandArea := 0

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			islandArea++

			x, y := cell.row, cell.col
			grid[x][y] = 0

			if x > 0 && grid[x-1][y] == 1 {
				queue = append(queue, Pos{x - 1, y})
				grid[x-1][y] = 0
			}
			if y > 0 && grid[x][y-1] == 1 {
				queue = append(queue, Pos{x, y - 1})
				grid[x][y-1] = 0
			}
			if x < rows-1 && grid[x+1][y] == 1 {
				queue = append(queue, Pos{x + 1, y})
				grid[x+1][y] = 0
			}
			if y < cols-1 && grid[x][y+1] == 1 {
				queue = append(queue, Pos{x, y + 1})
				grid[x][y+1] = 0
			}
		}
		maxArea = findMax(maxArea, islandArea)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	return maxArea
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
