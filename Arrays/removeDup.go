package arrays

/*
	LC #26
	TC--->O(n)
	SC--->O(1)
*/
func RemoveDuplicates(nums []int) int {
	pos := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
