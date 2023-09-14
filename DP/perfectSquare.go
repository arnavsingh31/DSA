package main

import (
	"math"
)

/*
LC #279, Same as Coin Change
TC--->O()TLE
*/
func perfectSquares(n int) int {
	squares := make([]int, 0)

	for i := 1; i*i < n; i++ {
		squares = append(squares, i*i)
	}

	return solveSquares(squares, n)
}

func solveSquares(arr []int, target int) int {
	if target == 0 {
		return 0
	}

	minCount := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] <= target {
			count := solveSquares(arr, target-arr[i])
			if count != math.MaxInt {
				minCount = min(1+count, minCount)
			}
		}
	}

	return minCount
}

/*
recursion + memoization
*/
func perfectSquares2(n int) int {
	squares := make([]int, 0)
	dp := make([]int, n+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	return solveSquaresMem(squares, n, &dp)
}

func solveSquaresMem(arr []int, target int, dp *[]int) int {
	if target == 0 {
		return 0
	}

	if (*dp)[target] != -1 {
		return (*dp)[target]
	}

	minCount := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] <= target {
			count := solveSquaresMem(arr, target-arr[i], dp)
			if count != math.MaxInt {
				minCount = min(minCount, 1+count)
			}
		}
	}

	(*dp)[target] = minCount

	return (*dp)[target]
}

/*
Tabulation method
TC--->O(n*len(squares))
SC--->O(n)
*/
func perfectSquares3(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < len(squares); j++ {
			if squares[j] <= i && dp[i-squares[j]] != math.MaxInt {
				dp[i] = min(1+dp[i-squares[j]], dp[i])
			}
		}
	}

	return dp[n]
}
