package main

/*	T.C -> O(n) and S.C --> O(n)
	make 2 auxillary arrays then which will we use to determine for any index what is the max height to the left
	of that index and max height to the right of the index.
	then we notice that water can only trap at an index if there it is a valley index i.e it has height greater
	to its left and right. Amount of max water that can be trapped will be equal to the
	= min(max_height to left of index, max_height to the right of index) - current height of the index.
*/
func trap(height []int) int {
	height_len := len(height)
	max_left := make([]int, height_len)
	max_right := make([]int, height_len)

	// populate max_left
	max_left[0] = height[0]
	for i := 1; i < height_len; i++ {
		max_left[i] = max_height(height[i], max_left[i-1])
	}

	max_right[height_len-1] = height[height_len-1]
	for i := height_len - 2; i >= 0; i-- {
		max_right[i] = max_height(height[i], max_right[i+1])
	}

	trapped_water := 0
	for i := 0; i < height_len; i++ {
		trapped_water += min_height(max_left[i], max_right[i]) - height[i]

	}

	return trapped_water
}

func min_height(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max_height(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// APPROAH #2 Using two pointers. T.C---> O(n), S.C-->O(1)
func trap2(height []int) int {
	height_len := len(height)

	// pointers for left and right elevations
	left_ptr, right_ptr := 0, height_len-1

	// max value of left and right elevation points
	max_left_ptr, max_right_ptr := height[left_ptr], height[right_ptr]
	trapped_water := 0

	// iterate until left elevation point reaches right elevation point or right elevation point reaches left elevation point.
	for left_ptr < right_ptr {
		// if max left elevation is smaller than max right elevation
		if max_left_ptr < max_right_ptr {
			left_ptr++                           // move the left elevation point by 1
			if height[left_ptr] > max_left_ptr { // if height of current left elevation after moving is greater than max left elevation
				max_left_ptr = height[left_ptr] // update the max left elevation with current left elevation point value
			}
			trapped_water += max_left_ptr - height[left_ptr] // store water with delta of max left elevation and current left elevation
		} else {
			right_ptr--                            // decrease the right elevation by 1
			if height[right_ptr] > max_right_ptr { // if height of current right elevation after decreasing is greater than max right elevation
				max_right_ptr = height[right_ptr] // update the max right elevation with current right elevation point value
			}
			trapped_water += max_right_ptr - height[right_ptr] // store water with delta of max right elevation and current right elevation
		}
	}
	return trapped_water
}
