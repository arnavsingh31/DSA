package matrices

/*
	LC #994, Multi-Source BFS
	TC--->O(m*n)
	SC--->O(m*n)
*/
func RottingOranges(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := make([]Pos, 0)
	freshOrange := 0

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	multiSourceBfs := func() int {
		time := 0

		for len(queue) > 0 {
			levelLength := len(queue)

			for i := 0; i < levelLength; i++ {
				cell := queue[0]
				queue = queue[1:]

				x, y := cell.row, cell.col

				for _, dir := range direction {
					newX, newY := x+dir[0], y+dir[1]

					if isCellValid(newX, newY) && grid[newX][newY] == 1 {
						queue = append(queue, Pos{newX, newY})
						grid[newX][newY] = 2
						freshOrange--
					}
				}
			}
			time++
			if freshOrange == 0 {
				return time
			}
		}
		return -1
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, Pos{i, j})
			}

			if grid[i][j] == 1 {
				freshOrange++
			}
		}
	}

	if freshOrange == 0 {
		return 0
	}

	return multiSourceBfs()
}
