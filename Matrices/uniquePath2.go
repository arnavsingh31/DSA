package main

/*
	LC #63
	TC--->O(m*n)
	SC--->O(m*n)
*/
func uniquePath2(grid [][]int) int {
	cache := make(map[Pos]int, 0)
	return dp2(0, 0, grid, &cache)
}

func dp2(i, j int, grid [][]int, cache *map[Pos]int) int {
	if val, exist := (*cache)[Pos{i, j}]; exist {
		return val
	}

	if i == len(grid) || j == len(grid[0]) || grid[i][j] == 1 {
		return 0
	}

	if i == len(grid)-1 && j == len(grid[0])-1 {
		return 1
	}

	right := dp2(i, j+1, grid, cache)
	down := dp2(i+1, j, grid, cache)
	(*cache)[Pos{i, j}] = right + down

	return right + down
}

/*
	TC--->O(m*n)
	SC--->O(n)-->columns
*/
func uniquePaths2_2(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dp := make([]int, cols)
	dp[0] = 1

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				if j > 0 {
					dp[j] += dp[j-1]
				}
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[cols-1]
}