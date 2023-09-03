package main

/*
	LC #417
	TC--->O(m*n)
	SC--->O(m*n)
*/

func waterFlow(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])
	pacificQueue := make([]Pos, 0)
	atlanticQueue := make([]Pos, 0)
	directions := [][]int{
		{-1, 0}, {0, 1}, {0, -1}, {1, 0},
	}
	ans := [][]int{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 || j == 0 {
				pacificQueue = append(pacificQueue, Pos{i, j})
			}

			if i == rows-1 || j == cols-1 {
				atlanticQueue = append(atlanticQueue, Pos{i, j})
			}
		}
	}

	isCellValid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < rows && j < cols
	}

	bfs := func(queue []Pos) [][]bool {
		visited := make([][]bool, rows)

		for i := 0; i < rows; i++ {
			visited[i] = make([]bool, cols)
		}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			x, y := cell.row, cell.col
			visited[x][y] = true

			for i := 0; i < len(directions); i++ {
				newX, newY := x+directions[i][0], y+directions[i][1]

				if isCellValid(newX, newY) && !visited[newX][newY] && heights[newX][newY] >= heights[x][y] {
					queue = append(queue, Pos{newX, newY})
					visited[newX][newY] = true
				}
			}
		}

		return visited
	}

	canReachPacific := bfs(pacificQueue)
	canReachAtlantic := bfs(atlanticQueue)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if canReachPacific[i][j] && canReachAtlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}
