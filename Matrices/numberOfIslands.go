package matrices

/*
	LC #200
	TC--->O(m*n)
	SC--->O(m*n)
*/

func NumIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	islands := 0

	isNeighbourLand := func(i, j int) bool {
		return grid[i][j] == '1'
	}

	bfs := func(i, j int) {
		queue := []Pos{{row: i, col: j}}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]

			grid[cell.row][cell.col] = '0'

			if cell.row > 0 && isNeighbourLand(cell.row-1, cell.col) {
				queue = append(queue, Pos{row: cell.row - 1, col: cell.col})
				grid[cell.row-1][cell.col] = '0'
			}

			if cell.col > 0 && isNeighbourLand(cell.row, cell.col-1) {
				queue = append(queue, Pos{row: cell.row, col: cell.col - 1})
				grid[cell.row][cell.col-1] = '0'
			}

			if cell.row < rows-1 && isNeighbourLand(cell.row+1, cell.col) {
				queue = append(queue, Pos{row: cell.row + 1, col: cell.col})
				grid[cell.row+1][cell.col] = '0'
			}

			if cell.col < cols-1 && isNeighbourLand(cell.row, cell.col+1) {
				queue = append(queue, Pos{row: cell.row, col: cell.col + 1})
				grid[cell.row][cell.col+1] = '0'
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				bfs(i, j)
				islands++
			}
		}
	}

	return islands
}

func NumIslands2(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make(map[Pos]bool, 0)
	islands := 0

	isCellLandAndNotVisited := func(i, j int) bool {
		return (grid[i][j] == '1') && !visited[Pos{row: i, col: j}]
	}

	bfs := func(i, j int) {
		queue := []Pos{{row: i, col: j}}

		for len(queue) > 0 {
			cell := queue[0]
			queue = queue[1:]
			visited[cell] = true

			// mark neighbour visited to avoid multiple addition of same cells in the queue, which leads to out of memory error.
			if cell.row > 0 && isCellLandAndNotVisited(cell.row-1, cell.col) {
				neighbour := Pos{row: cell.row - 1, col: cell.col}
				queue = append(queue, neighbour)
				visited[neighbour] = true
			}

			if cell.col > 0 && isCellLandAndNotVisited(cell.row, cell.col-1) {
				neighbour := Pos{row: cell.row, col: cell.col - 1}
				queue = append(queue, Pos{row: cell.row, col: cell.col - 1})
				visited[neighbour] = true
			}

			if cell.row < rows-1 && isCellLandAndNotVisited(cell.row+1, cell.col) {
				neighbour := Pos{row: cell.row + 1, col: cell.col}
				queue = append(queue, Pos{row: cell.row + 1, col: cell.col})
				visited[neighbour] = true
			}

			if cell.col < cols-1 && isCellLandAndNotVisited(cell.row, cell.col+1) {
				neighbour := Pos{row: cell.row, col: cell.col + 1}
				queue = append(queue, Pos{row: cell.row, col: cell.col + 1})
				visited[neighbour] = true
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[Pos{row: i, col: j}] {
				bfs(i, j)
				islands++
			}
		}
	}

	return islands
}

/*
Testcase
[
{"1","1","1","1","1","0","1","1","1","1","1","1","1","1","1","0","1","0","1","1"},
{"0","1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","0"},
{"1","0","1","1","1","0","0","1","1","0","1","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","0","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"1","0","0","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"1","0","1","1","1","1","1","1","0","1","1","1","0","1","1","1","0","1","1","1"},
{"0","1","1","1","1","1","1","1","1","1","1","1","0","1","1","0","1","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","0","1","1"},
{"1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"0","1","1","1","1","1","1","1","0","1","1","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","1","0","1","1","1","1","1","1","1","0","1","1","1","1","1","1"},
{"1","0","1","1","1","1","1","0","1","1","1","0","1","1","1","1","0","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","1","0"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","0","0"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"},
{"1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"}
]
*/
