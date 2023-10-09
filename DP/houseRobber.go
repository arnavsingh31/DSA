package dp

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #198. using dp, space optimised solution. For all solutioncheck similar concept problem
	maxSumOfNonAdjacentElement
	TC--->O(n)
	SC--->O(1)
*/
func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prev2 := nums[0]
	prev1 := util.Max(prev2, nums[1])

	for i := 2; i < len(nums); i++ {
		rob := nums[i] + prev2
		dontRob := prev1
		maxMoney := util.Max(rob, dontRob)
		prev2 = prev1
		prev1 = maxMoney
	}

	return prev1
}
