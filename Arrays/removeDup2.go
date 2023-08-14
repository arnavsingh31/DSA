package main

/*
	LC #80
	TC--->O(n)
	SC--->O(1)
*/
func removeDup2(nums []int) int {
	pos := 1
	i := 1
	count := 1

	for i < len(nums) {
		if nums[i] == nums[i-1] && count < 2 {
			nums[pos] = nums[i]
			pos++
			count++
		}

		if nums[i] != nums[i-1] {
			count = 1
			nums[pos] = nums[i]
			pos++
		}
		i++
	}
	return pos
}
