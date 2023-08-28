package main

/*
	LC #62
	TC--->O(m*n)
	SC--->O(m*n) call stack + cache
*/
func uniquePath(m int, n int) int {
	cache := make(map[Pos]int, 0)
	return dpUniquePath(0, 0, m, n, &cache)
}

func dpUniquePath(i, j, m, n int, cache *map[Pos]int) int {
	if val, exist := (*cache)[Pos{i, j}]; exist {
		return val
	}

	if i == m {
		return 0
	}
	if j == n {
		return 0
	}

	if i == m-1 && j == n-1 {
		return 1
	}

	right := dpUniquePath(i, j+1, m, n, cache)
	down := dpUniquePath(i+1, j, m, n, cache)
	(*cache)[Pos{i, j}] = right + down

	return right + down
}

/*
	TC-->O(m*n)
	SC-->O(m*n) only space required for grid
*/
func uniquePath_2(m, n int) int {
	grid := make([][]int, m)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				grid[i][j] += grid[i-1][j] + grid[i][j-1]
			} else if i > 0 {
				grid[i][j] += grid[i-1][j]
			} else if j > 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[m-1][n-1]
}
