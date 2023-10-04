package matrices

import "sort"

/*
	LC #2500
	TC--->O(mlogm + m*n), m-> col and n->rows. mlogm for sorting each row of m elements + m*n since we are
	traversing every element of grid once.
	SC--->O(1)
*/
func DeleteGreatest(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	sum := 0

	for _, arr := range grid {
		sort.Ints(arr)
	}

	for j := cols - 1; j >= 0; j-- {
		max := 0
		for i := 0; i < rows; i++ {
			if max < grid[i][j] {
				max = grid[i][j]
			}
		}
		sum += max
	}

	return sum
}
