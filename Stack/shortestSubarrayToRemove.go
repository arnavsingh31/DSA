package stack

import util "github.com/arnavsingh31/DSA/Util"

/*
	LC #1574
	Explanation:https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/solutions/830480/c-o-n-sliding-window-explanation-with-illustrations/
*/
func FindLengthOfSubArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < len(arr)-1 && arr[left] <= arr[left+1] {
		left++
	}
	if left == len(arr)-1 {
		return 0
	}

	for right > left && arr[right] >= arr[right-1] {
		right--
	}
	res := util.Min(right, len(arr)-left-1)

	i := 0
	j := right
	for i <= left && j <= len(arr)-1 {
		if arr[j] >= arr[i] {
			res = util.Min(res, j-i-1)
			i++
		} else {
			j++
		}
	}

	return res
}
