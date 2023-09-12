package main

/*
	TC---> exponential {TLE} recursion only,
	TC---> O(n)recursion+ memoization {solvable}, SC--->O(n){call stack + dp}
*/
func maxSum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	dp := make([]int, len(arr)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	return solve(arr, 0, &dp)
}

func solve(arr []int, i int, dp *[]int) int {
	if i >= len(arr) {
		return 0
	}

	if (*dp)[i] != 0 {
		return (*dp)[i]
	}

	include := arr[i] + solve(arr, i+2, dp)
	exclude := solve(arr, i+1, dp)

	(*dp)[i] = max(include, exclude)
	return (*dp)[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
	tabuation method
	TC-->O(n)
	SC--->O(n)
*/

func maxSum2(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = max(dp[0], arr[1])

	for i := 2; i < len(arr); i++ {
		incl := arr[i] + dp[i-2]
		exc := dp[i-1]
		dp[i] = max(incl, exc)
	}

	return dp[len(arr)-1]
}

/*
	we can optimize tabulation solution
	TC--->O(n)
	SC--->O(1)
*/
func maxSum3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prev2 := nums[0]
	prev1 := max(prev2, nums[1])

	for i := 2; i < len(nums); i++ {
		include := nums[i] + prev2
		exclude := prev1
		maxMoney := max(include, exclude)
		prev2 = prev1
		prev1 = maxMoney
	}

	return prev1
}
