package matrices

type Pos struct {
	row int
	col int
}

/*
LC #289
TC--->O(m*n)
SC--->O(m*n)
*/
func GameOfLife(board [][]int) {
	rows := len(board)
	cols := len(board[0])
	currenStateMap := make(map[Pos]int, 0)

	isCellAlive := func(i, j int) bool {
		if val, exist := currenStateMap[Pos{row: i, col: j}]; exist {
			return val == 1
		}
		return board[i][j] == 1
	}

	neighbourLiveCount := func(i, j int) int {
		aliveNeighbours := 0

		if i > 0 && isCellAlive(i-1, j) {
			aliveNeighbours++
		}
		if j > 0 && isCellAlive(i, j-1) {
			aliveNeighbours++
		}
		if i < rows-1 && isCellAlive(i+1, j) {
			aliveNeighbours++
		}
		if j < cols-1 && isCellAlive(i, j+1) {
			aliveNeighbours++
		}
		if i > 0 && j > 0 && isCellAlive(i-1, j-1) {
			aliveNeighbours++
		}
		if i > 0 && j < cols-1 && isCellAlive(i-1, j+1) {
			aliveNeighbours++
		}
		if i < rows-1 && j > 0 && isCellAlive(i+1, j-1) {
			aliveNeighbours++
		}
		if i < rows-1 && j < cols-1 && isCellAlive(i+1, j+1) {
			aliveNeighbours++
		}

		return aliveNeighbours
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// populate current state of cell.
			pos := Pos{
				row: i,
				col: j,
			}
			currenStateMap[pos] = board[i][j]

			// update board state
			aliveNeighbour := neighbourLiveCount(i, j)

			if board[i][j] == 1 && (aliveNeighbour < 2 || aliveNeighbour > 3) {
				board[i][j] = 0
			} else if board[i][j] == 0 && aliveNeighbour == 3 {
				board[i][j] = 1
			}
		}
	}

}

/*
	initial  |  new  |  map
		0		 0		 0
		0		 1		 2
		1	     0		 3
		1		 1		 1

	TC--->O(m*n)
	SC--->O(1)

*/

func GameOfLife2(board [][]int) {
	rows := len(board)
	cols := len(board[0])

	isCellAlive := func(i, j int) bool {
		cell := board[i][j]

		if cell == 1 || cell == 3 {
			return true
		}
		return false
	}

	neighbourLiveCount := func(i, j int) int {
		aliveNeighbours := 0

		if i > 0 && isCellAlive(i-1, j) {
			aliveNeighbours++
		}
		if j > 0 && isCellAlive(i, j-1) {
			aliveNeighbours++
		}
		if i < rows-1 && isCellAlive(i+1, j) {
			aliveNeighbours++
		}
		if j < cols-1 && isCellAlive(i, j+1) {
			aliveNeighbours++
		}
		if i > 0 && j > 0 && isCellAlive(i-1, j-1) {
			aliveNeighbours++
		}
		if i > 0 && j < cols-1 && isCellAlive(i-1, j+1) {
			aliveNeighbours++
		}
		if i < rows-1 && j > 0 && isCellAlive(i+1, j-1) {
			aliveNeighbours++
		}
		if i < rows-1 && j < cols-1 && isCellAlive(i+1, j+1) {
			aliveNeighbours++
		}

		return aliveNeighbours
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aliveNeighbour := neighbourLiveCount(i, j)

			if board[i][j] == 1 && (aliveNeighbour < 2 || aliveNeighbour > 3) {
				board[i][j] = 3
			} else if board[i][j] == 0 && aliveNeighbour == 3 {
				board[i][j] = 2
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			}

			if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}
