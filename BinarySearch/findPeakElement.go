package binarysearch

/*
	LC #162
	TC--->O(n)
	SC--->O(1)
*/
func FindPeakElement(nums []int) int {
	n := len(nums)

	// handling edge cases
	if n == 1 || nums[0] > nums[1] {
		return 0
	}

	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	return -1
}

/*
	TC--->O(logn)
	SC--->O(1)
	Binary Search approach.
*/
func FindPeakElement2(nums []int) int {
	n := len(nums)

	// handling edge cases
	if n == 1 || nums[0] > nums[1] {
		return 0
	}

	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	left, right := 1, n-2
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1

}
