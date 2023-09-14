package main

/*
	GFG Problem. {Watch video for concept}
	Count Derangements (Permutation such that no element appears in its original position).
	A Derangement is a permutation of n elements, such that no element appears in its original position. For example, a derangement of {0, 1, 2, 3} is {2, 3, 1, 0}.
	Given a number n, find the total number of Derangements of a set of n elements.
	TLE
	TC--> exponential
*/
func countDerangements(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	ans := (n - 1) * (countDerangements(n-1) + countDerangements(n-2))

	return ans
}

/*
	Recursion + Memoization
	TC--->O(n)
	SC--->O(n)+O(n){call stack + dp}
*/
func countDerangements2(n int) int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	return solveDerangements(n, &dp)
}

func solveDerangements(n int, dp *[]int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if (*dp)[n] != -1 {
		return (*dp)[n]
	}

	(*dp)[n] = (n - 1) * (solveDerangements(n-1, dp) + solveDerangements(n-2, dp))

	return (*dp)[n]
}

/*
	Tabulation method
	TC--->O(n)
	SC--->O(n)
*/
func countDerangements3(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = (i - 1) * (dp[i-1] + dp[i-2])
	}

	return dp[n]
}

/*
	Space Optimisation
	TC--->O(n)
	SC--->O(1)
*/
func countDerangements4(n int) int {
	x1, x2 := 0, 1

	for i := 2; i <= n; i++ {
		ans := (i - 1) * (x2 + x1)
		x1 = x2
		x2 = ans
	}
	return x2

}
