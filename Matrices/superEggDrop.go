package main

import (
	"math"
)

/*
LC #887 also covers the problem "egg drop with 2 eggs and N floors".
Return minimum number of attempts need to find the high enough floor from which egg breaks
{TLE}
*/
func eggDrop(n, k int) int {
	// no more floors to check
	if n == 0 || n == 1 {
		return n
	}

	// only 1 egg remains so we have to check linearly {worst case would be to check all remaining n floors}
	if k == 1 {
		return n
	}

	dropCount := math.MaxInt
	for i := 1; i <= n; i++ {
		res := max(eggDrop(i-1, k-1), eggDrop(n-i, k))
		dropCount = min(dropCount, res)
	}

	return dropCount + 1

}

/*
Recursion + Memoization
TC--->O(n*k*k)
SC--->O(n*k)
TLE :(
*/
func eggDrop2(n, k int) int {
	dp := make([][]int, 1000)

	for i := 0; i < 1000; i++ {
		dp[i] = make([]int, 1000)
	}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			dp[i][j] = -1
		}
	}

	return solveMem(n, k, &dp)
}

func solveMem(n, k int, dp *[][]int) int {
	if (*dp)[n][k] != -1 {
		return (*dp)[n][k]
	}

	if n == 0 || n == 1 {
		return n
	}

	if k == 1 {
		return n
	}

	count := math.MaxInt
	for i := 1; i <= n; i++ {
		res := max(solveMem(i-1, k-1, dp), solveMem(n-i, k, dp))
		count = min(count, res)
	}

	(*dp)[n][k] = 1 + count
	return (*dp)[n][k]
}

/*
	Tabulation Method
	Prefilled matrix:
	[ k{egg}->rows, n{floors}->col
		[0 0 0 0 0 0 0 0 0]
		[0 1 2 3 4 5 6 7 8]
		[0 1 0 0 0 0 0 0 0]
	]
	TC--->O(n*k*k)
	SC--->O(n*k)
	Also gives TLE. :(
*/

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)

	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}

	for j := 1; j <= n; j++ {
		dp[1][j] = j
	}

	for i := 1; i <= k; i++ {
		dp[i][1] = 1
		dp[i][0] = 0
	}

	for i := 2; i <= k; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = math.MaxInt

			for x := 1; x <= j; x++ {
				res := 1 + max(dp[i-1][x-1], dp[i][j-x])
				dp[i][j] = min(res, dp[i][j])
			}
		}
	}

	return dp[k][n]
}
