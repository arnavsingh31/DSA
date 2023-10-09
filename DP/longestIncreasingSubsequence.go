package dp

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #300
*/
func LengthOfLIS(nums []int) int {
	dp := make([][]int, len(nums))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums)+1)
		for j := 0; j < len(nums)+1; j++ {
			dp[i][j] = -1
		}
	}

	return solveRec(nums, 0, -1, &dp)
}

func solveRec(nums []int, index int, prevIndex int, dp *[][]int) int {
	//base case
	if index == len(nums) {
		return 0
	}

	if (*dp)[index][prevIndex+1] != -1 {
		return (*dp)[index][prevIndex+1]
	}

	// recurrence
	// don't take current element in our subsequence
	notTake := 0 + solveRec(nums, index+1, prevIndex, dp)

	// take current element in our subsequence if----
	take := 0
	if prevIndex == -1 || nums[index] > nums[prevIndex] {
		take = 1 + solveRec(nums, index+1, index, dp)
	}

	(*dp)[index][prevIndex+1] = util.Max(take, notTake)

	return (*dp)[index][prevIndex+1]
}
