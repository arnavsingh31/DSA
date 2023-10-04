package matrices

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #221
	TC--->O(m*n)
	SC--->O(m*n)
	2D-DP memoization + recursion
*/
func MaximalSquareArea(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	maxSide := 0
	dp := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dp[i][j] = -1
		}
	}

	var findSquare func(int, int) int
	findSquare = func(i, j int) int {
		if i >= rows || j >= cols {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		right := findSquare(i, j+1)
		diagonal := findSquare(i+1, j+1)
		down := findSquare(i+1, j)

		if matrix[i][j] == '1' {
			min := 1 + util.Min(right, util.Min(diagonal, down))
			maxSide = util.Max(maxSide, min)
			dp[i][j] = min
			return min
		} else {
			dp[i][j] = 0
			return 0
		}
	}

	findSquare(0, 0)
	return maxSide * maxSide
}

/*
	2d- DP, Tabulation method
*/
func MaximalSquareArea2(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows)
	maxSide := 0

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					top := dp[i-1][j]
					left := dp[i][j-1]
					diagonal := dp[i-1][j-1]

					dp[i][j] = 1 + util.Min(top, util.Min(left, diagonal))
				}
				maxSide = util.Max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}
