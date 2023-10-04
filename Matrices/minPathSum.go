package matrices

import (
	"math"

	util "github.com/arnavsingh31/DSA/Util"
)

/*
LC #64
TC--->O(m*n)
SC--->O(m*n), call stack + memo
Upon visualisation this problem is like finding minPathSum in a binary tree, where we can apply dfs algo.
But alone dfs is not sufficient so we use memoization technique to avoid TLE.(DP)
FYI

	type Pos struct {
		row int
		col int
	}
*/
func MinPathSum(matrix [][]int) int {
	pathSumMap := make(map[Pos]int, 0)
	return dp(0, 0, matrix, &pathSumMap)
}

func dp(i, j int, matrix [][]int, mem *map[Pos]int) int {
	if val, exist := (*mem)[Pos{i, j}]; exist {
		return val
	}

	if i == len(matrix) {
		return math.MaxInt
	}

	if j == len(matrix[0]) {
		return math.MaxInt
	}

	if i == len(matrix)-1 && j == len(matrix[0])-1 {
		return matrix[i][j]
	}

	right := dp(i, j+1, matrix, mem)
	down := dp(i+1, j, matrix, mem)

	(*mem)[Pos{row: i, col: j}] = matrix[i][j] + util.Min(right, down)

	return matrix[i][j] + util.Min(right, down)
}

/*
	 1
	/  \
	1   3
   /  \
   4    5
    \  /  \
	 2 2   1
	  \ \  /
	   1 1 1
*/

/*
TC--->O(m*n)
SC--->O(1), Optimised sol.
*/
func MinPathSum2(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i > 0 && j > 0 {
				matrix[i][j] += util.Min(matrix[i][j-1], matrix[i-1][j])
			} else if i > 0 {
				matrix[i][j] += matrix[i-1][j]
			} else if j > 0 {
				matrix[i][j] += matrix[i][j-1]
			}
		}
	}

	return matrix[rows-1][cols-1]
}
