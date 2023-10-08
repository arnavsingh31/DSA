package arrays

/*
	LC #27
	TC--->O(n)
	SC--->O(1)
*/
func RemoveElements(nums []int, val int) int {
	pos := 0

	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
