package matrices

/*
	LC #463
	TC--->O(m*n)
	SC--->O(1)
*/
func IslandPerimeter(grid [][]int) int {
	perimeter, rows, cols := 0, len(grid), len(grid[0])

	isLand := func(i, j int) bool {
		return grid[i][j] == 1
	}

	neighbourIsland := func(i, j int) int {
		count := 0

		if i > 0 && isLand(i-1, j) {
			count++
		}
		if j > 0 && isLand(i, j-1) {
			count++
		}
		if i < rows-1 && isLand(i+1, j) {
			count++
		}
		if j < cols-1 && isLand(i, j+1) {
			count++
		}
		return count
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				perimeter += 4 - neighbourIsland(i, j)
			}
		}
	}

	return perimeter
}
