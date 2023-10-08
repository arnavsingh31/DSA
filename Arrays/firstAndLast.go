package arrays

// T.C--> O(n)
func FirstAndLastIndex(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ans := []int{}
	isFirst := true
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && isFirst {
			ans = append(ans, i)
			isFirst = false
			count++
		} else if nums[i] == target {
			count++
		}
	}

	if len(ans) > 0 {
		ans = append(ans, ans[0]+count-1)
	} else {
		ans = append(ans, -1, -1)
	}

	return ans
}

// T.C---> O(logn) using 2 binary search
func FirstAndLastIndex2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ans := []int{}
	low, high := 0, len(nums)-1
	first_pos := -1
	// first find the 1st occurrence of target using binary search
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			first_pos = mid
			high = mid - 1
		} else if nums[mid] > target {
			high = mid - 1
		} else { // nums[mid] <  target
			low = mid + 1
		}
	}
	last_pos := -1
	low2, high2 := 0, len(nums)-1
	for low2 <= high2 {
		mid := low2 + (high2-low2)/2

		if nums[mid] == target {
			last_pos = mid
			low2 = mid + 1
		} else if nums[mid] > target {
			high2 = mid - 1
		} else {
			low2 = mid + 1
		}
	}
	ans = append(ans, first_pos, last_pos)

	return ans
}

// Approach 3 is improved version of binary search used in above approach. here we find first positon of
// greaterThanOrEqual(>=target)---> this give us first index of target and (>=target+1)---> this gives us
// first index of target+1, then we decrement it by 1 to get the last index of target.
func FirstAndLastIndex3(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first_pos := greaterThanEqualPos(nums, target)
	last_pos := greaterThanEqualPos(nums, target+1) - 1

	if first_pos <= last_pos {
		return []int{first_pos, last_pos}
	}
	return []int{-1, -1}
}

func greaterThanEqualPos(arr []int, target int) int {
	n := len(arr)
	low := 0
	high := n - 1
	first_pos := n

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] >= target {
			first_pos = mid
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		}
	}

	return first_pos
}
