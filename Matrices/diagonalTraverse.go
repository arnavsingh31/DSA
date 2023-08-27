package main

/*
	LC #498 Same as zig-zag traversal of binary tree. (BFS level wise)
	TC--->O(m*n)
	SC--->O(m*n)
*/

func diagonalTraverse(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	visited := make(map[Pos]bool, 0)
	ans := []int{}
	reverse := false

	isCellVisited := func(i, j int) bool {
		return visited[Pos{row: i, col: j}]
	}

	bfs := func(i, j int) {
		queue := []Pos{{row: i, col: j}}

		for len(queue) > 0 {
			levelLength := len(queue)
			arr := make([]int, 0)

			for i := 0; i < levelLength; i++ {
				cell := queue[0]
				queue = queue[1:]
				visited[cell] = true

				if reverse {
					arr = append(append([]int{}, matrix[cell.row][cell.col]), arr...)
				} else {
					arr = append(arr, matrix[cell.row][cell.col])
				}

				if cell.row < rows-1 && !isCellVisited(cell.row+1, cell.col) {
					neighbour := Pos{row: cell.row + 1, col: cell.col}
					queue = append(queue, neighbour)
					visited[neighbour] = true
				}

				if cell.col < cols-1 && !isCellVisited(cell.row, cell.col+1) {
					neighbour := Pos{row: cell.row, col: cell.col + 1}
					queue = append(queue, neighbour)
					visited[neighbour] = true
				}
			}

			reverse = !reverse
			ans = append(ans, arr...)

		}
	}

	bfs(0, 0)
	return ans
}

/*
	TC--->O(m*n)
	SC--->O(1)

	We need to figure out two things for each diagonal:

	- The direction in which we want to process it's elements and
	- The head or the starting point for the diagonal depending upon its direction.
	- The direction is pretty straightforward. We can simply use a boolean variable and keep alternating it to figure out the direction for a diagonal.

	Next head when going UP
	- The head would be the node directly below the tail of the previous diagonal. Unless the tail lies in the last row of the matrix in which case the head would be the node right next to the tail.

	Next head when going DOWN
	- The head would be the node to the right of the tail of the previous diagonal. Unless the tail lies in the last column of the matrix in which case the head would be the node directly below the tail.

*/
func diagonalTraverse2(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	ans := []int{}
	direction := true // true-> going to upward, false -> going downward
	i, j := 0, 0

	for len(ans) < rows*cols {
		ans = append(ans, matrix[i][j])

		if direction {
			if i == 0 && j < cols-1 {
				// reached end of up-diagonal, reversing direction
				direction = !direction
				// next head
				j += 1
			} else if j == cols-1 {
				// reached end of up-diagonal, reversing direction
				direction = !direction
				// next head
				i += 1
			} else {
				// moving pointer to next element along up-diagonal
				i -= 1
				j += 1
			}
		} else {
			if j == 0 && i < rows-1 {
				// reached end of down-diagonal, reversing direction
				direction = !direction
				i += 1
			} else if i == rows-1 {
				// reached end of down-diagonal, reversing direction
				direction = !direction
				// next head
				j += 1
			} else {
				// moving pointer to next element along down-diagonal
				i += 1
				j -= 1
			}
		}
	}

	return ans
}
