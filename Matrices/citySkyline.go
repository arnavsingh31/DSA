package main

/*
	LC #807
	TC--->O(n*n)
	SC--->O(n+n)
*/
func maxIncreaseKeepingSkyline(grid [][]int) int {
	// for each cell(i,j) we need max val in ith row and max val in jth col. then we modify height
	// of current cell by newHeight = min(rowMax, colMax)
	rows := len(grid)
	cols := len(grid[0])
	rowMaxArr := make([]int, rows)
	colMaxArr := make([]int, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rowMaxArr[i] = max(rowMaxArr[i], grid[i][j])
			colMaxArr[j] = max(colMaxArr[j], grid[i][j])
		}
	}

	increasedSum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			prev := grid[i][j]
			grid[i][j] = min(rowMaxArr[i], colMaxArr[j])
			increasedSum += grid[i][j] - prev
		}
	}

	return increasedSum
}
