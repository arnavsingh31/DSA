package matrices

/*
	LC #934
	TC--->O(n*n)
	SC--->O(n*n)
	First apply BFS to find all cells of 1st island and mark them to 2 (in-order to distinguish
	from island 2 cells). Now use these 1st island cells and apply bfs again. If we find cell with
	val == 1 then we encountered island2 so return distance. But if we encounter water (cell == 0)
	then add it to new Queue. Once iteration ends, set secondBfsQueue = newQueue, and inc distance
	by 1. repeat bfs on secondBfsQueue(water cells) till we reach island2.
*/
func ShortestBridge(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	secondBfsQueue := []Pos{}
	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	bfs := func(i, j int) {
		queue := []Pos{{i, j}}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			x, y := cell.row, cell.col
			secondBfsQueue = append(secondBfsQueue, Pos{x, y})
			grid[x][y] = 2

			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]

				if isCellValid(newX, newY) && grid[newX][newY] == 1 {
					queue = append(queue, Pos{newX, newY})
					grid[newX][newY] = 2
				}
			}
		}
	}

	for i := 0; i < rows; i++ {
		firstIslandExplored := false
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
				firstIslandExplored = true
				break
			}
		}

		if firstIslandExplored {
			break
		}
	}

	distance := 0
	for len(secondBfsQueue) > 0 {
		waterCells := []Pos{} // new queue

		for _, cell := range secondBfsQueue {
			x, y := cell.row, cell.col
			grid[x][y] = -1

			for _, dir := range directions {
				newX, newY := x+dir[0], y+dir[1]

				if isCellValid(newX, newY) {
					if grid[newX][newY] == 1 {
						return distance
					} else if grid[newX][newY] == 0 {
						waterCells = append(waterCells, Pos{newX, newY})
						grid[newX][newY] = -1
					}
				}
			}
		}

		secondBfsQueue = waterCells
		distance++
	}

	return distance
}
