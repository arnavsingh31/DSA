package main

import "math"

/*
	{GFG question}
	Given a rod of length L, the task is to cut the rod in such a way that the total number of segments of length x, y, and z is maximized. The segments can only be of length x, y, and z.

	Recurive sol {TLE}
*/
func cutRod(n, x, y, z int) int {
	segments := solveRodRec(n, x, y, z)

	if segments < 0 {
		return 0
	}

	return segments
}

func solveRodRec(n, x, y, z int) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return math.MinInt
	}

	a := 1 + solveRodRec(n-x, x, y, z)
	b := 1 + solveRodRec(n-y, x, y, z)
	c := 1 + solveRodRec(n-y, x, y, z)

	ans := max(a, max(b, c))

	return ans
}

/*
	Recursion + Memoization
*/
func cutRod2(n, x, y, z int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	segments := solveRodMem(n, x, y, z, &dp)

	if segments < 0 {
		return 0
	}

	return segments
}

func solveRodMem(n, x, y, z int, dp *[]int) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return math.MinInt
	}

	if (*dp)[n] != -1 {
		return (*dp)[n]
	}

	a := 1 + solveRodMem(n-x, x, y, z, dp)
	b := 1 + solveRodMem(n-y, x, y, z, dp)
	c := 1 + solveRodMem(n-y, x, y, z, dp)

	(*dp)[n] = max(a, max(b, c))

	return (*dp)[n]
}

/*
	Tabulation method
*/
func cutRod3(n, x, y, z int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = math.MinInt
	}

	dp[0] = 0

	for i := 1; i <= n; i++ {
		if i-x >= 0 {
			dp[i] = max(1+dp[i-x], dp[i])
		}
		if i-y >= 0 {
			dp[i] = max(1+dp[i-y], dp[i])
		}

		if i-z >= 0 {
			dp[i] = max(1+dp[i-z], dp[i])
		}
	}

	if dp[n] < 0 {
		return 0
	}

	return dp[n]
}
