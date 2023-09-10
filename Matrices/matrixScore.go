package main

import (
	"math"
)

/*
LC #861
TC--->O(m*n)
SC--->O(m+n)
*/
func matrixScore(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	toggleRowArr := make([]int, rows)
	toggleColArr := make([]int, cols)
	// first toggle row wise and then toggle col wise
	// for row if first element is 0 then toggle entire row else keep it as it is
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][0] == 0 {
				toggleRowArr[i] = 1
			}
		}
	}

	// for toggle use XOR with 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// toggle row
			if toggleRowArr[i] == 1 {
				matrix[i][j] ^= 1
			}
		}
	}

	// for col if zero count is > num 1 (aka number of rows/2) then toggle col
	for j := 0; j < cols; j++ {
		zeroCount := 0
		for i := 0; i < rows; i++ {
			if matrix[i][j] == 0 {
				zeroCount++
			}
		}
		if zeroCount > rows/2 {
			toggleColArr[j] = 1
		}
	}

	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			// toggle col
			if toggleColArr[j] == 1 {
				matrix[i][j] ^= 1
			}
		}
	}

	// find sum
	sum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				sum += int(math.Pow(2, float64(cols-1-j)))
			}
		}
	}
	return sum
}
