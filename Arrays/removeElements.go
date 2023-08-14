package main

/*
	LC #27
	TC--->O(n)
	SC--->O(1)
*/
func removeElements(nums []int, val int) int {
	pos := 0

	for _, num := range nums {
		if num != val {
			nums[pos] = num
			pos++
		}
	}
	return pos
}
