package dp

import util "github.com/arnavsingh31/DSA/Util"

/*
	TC---> exponential {TLE} recursion only,
	TC---> O(n)recursion+ memoization {solvable}, SC--->O(n){call stack + dp}
*/
func MaxSum(arr []int) int {
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

	(*dp)[i] = util.Max(include, exclude)
	return (*dp)[i]
}

/*
	tabuation method
	TC-->O(n)
	SC--->O(n)
*/

func MaxSum2(arr []int) int {
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	dp[1] = util.Max(dp[0], arr[1])

	for i := 2; i < len(arr); i++ {
		incl := arr[i] + dp[i-2]
		exc := dp[i-1]
		dp[i] = util.Max(incl, exc)
	}

	return dp[len(arr)-1]
}

/*
	we can optimize tabulation solution
	TC--->O(n)
	SC--->O(1)
*/
func MaxSum3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prev2 := nums[0]
	prev1 := util.Max(prev2, nums[1])

	for i := 2; i < len(nums); i++ {
		include := nums[i] + prev2
		exclude := prev1
		maxMoney := util.Max(include, exclude)
		prev2 = prev1
		prev1 = maxMoney
	}

	return prev1
}
