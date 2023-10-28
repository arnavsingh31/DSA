package twopointer

// LC #287
/*
	we use fast and slow algorithm to determine the cycle in the nums array. After finding cycle we stop slow
	pointer and reset fast pointer to nums[0] and move both pointer at same pace. The point where they meet
	is the beginning of cycle in nums/our duplicate number.
*/
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast

}
