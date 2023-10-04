package matrices

/*
	LC #827
	TC--->O(n^4) {TLE}
	SC--->O(n^2) {queue + map}
*/
func LargestIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	maxArea := 0
	hasZero := false

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	bfs := func(i, j int) {
		queue := []Pos{{i, j}}
		area := 0
		visited := make(map[Pos]bool)

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			area++

			x, y := cell.row, cell.col
			visited[cell] = true

			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]

				if isCellValid(newX, newY) && grid[newX][newY] == 1 && !visited[Pos{newX, newY}] {
					queue = append(queue, Pos{newX, newY})
					visited[Pos{newX, newY}] = true
				}
			}
		}

		if area > maxArea {
			maxArea = area
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				hasZero = true
				grid[i][j] = 1
				bfs(i, j)
				grid[i][j] = 0
			}
		}
	}

	if !hasZero {
		return rows * cols
	}

	return maxArea
}

/*
	Optimised solution using map.
	TC--->O(n^n)
	SC--->O(n^n)
*/

func LargestIsland2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	maxArea := 0
	regionId := 2
	regionAreaMap := make(map[int]int)
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	bfs := func(i, j int) {
		queue := []Pos{{i, j}}
		area := 0

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			area++

			x, y := cell.row, cell.col
			grid[x][y] = regionId

			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]

				if isCellValid(newX, newY) && grid[newX][newY] == 1 {
					queue = append(queue, Pos{newX, newY})
					grid[newX][newY] = regionId
				}
			}
		}

		regionAreaMap[regionId] = area

		if area > maxArea {
			maxArea = area
		}

		regionId++
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				idVisited := make(map[int]bool)
				area := 1

				for _, dir := range directions {
					x, y := i+dir[0], j+dir[1]

					if isCellValid(x, y) && grid[x][y] > 1 && !idVisited[grid[x][y]] {
						area += regionAreaMap[grid[x][y]]
						idVisited[grid[x][y]] = true
					}
				}

				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
