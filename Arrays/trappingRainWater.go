package main

/*	T.C -> O(n) and S.C --> O(n) [More intuitive than two pointer approach].
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

/*
	APPROAH #2 Using two pointers. T.C---> O(n), S.C-->O(1)

	Intuition
	The intuition behind this algorithm is to use two pointers to scan the height slice from both ends. By focusing on the lower height at each step, we can determine the potential trapped water. We keep track of the maximum height seen on each side and calculate the trapped water based on these maximums. The algorithm efficiently calculates the trapped water without using additional space and runs in linear time, making it an effective solution to the problem.

	Approach

	1. Initialization: We start with two pointers, left and right, initialized to the first and last positions of the height slice, respectively. We also initialize two variables, leftMax and rightMax, to keep track of the maximum height seen so far on the left and right sides, respectively. Lastly, we initialize a variable result to store the total trapped water, initially set to zero.

	2. Processing: The two pointers (left and right) move towards each other until they meet. During this process, we focus on the lower of the two heights at the current positions, as this determines the potential trapped water.

	3. Trapping Water: When the height at the left pointer is smaller than the height at the right pointer, we process the left side. If the current height at the left pointer is greater than or equal to the leftMax, we update leftMax to this current height. If the current height is smaller than leftMax, it means there is a potential trap, so we add the difference between leftMax and the current height to the result, which represents the trapped water at this position. After this, we move the left pointer to the right.

	4. Similarly, when the height at the right pointer is smaller than or equal to the height at the left pointer, we process the right side. We update rightMax or add trapped water to the result as necessary. Then, we move the right pointer to the left.

	Continue the process until left and right meet. At this point, we have calculated the total trapped water, which is stored in the result variable.
*/
func trap2(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	result := 0

	// Move the left and right pointers towards each other until they meet.
	for left < right {
		// When height at left pointer is smaller than height at right pointer,
		// we process the left side.
		if height[left] < height[right] {
			// If the current height at left is greater than or equal to the leftMax,
			// update the leftMax with the current height.
			// Otherwise, add the difference between the leftMax and current height to the result,
			// which represents the trapped water at the current position.
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			// Move the left pointer to the right.
			left++
		} else {
			// When height at right pointer is smaller than or equal to height at left pointer,
			// we process the right side.
			// Similar to the left side, update rightMax or add trapped water to the result.
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			// Move the right pointer to the left.
			right--
		}
	}
	return result
}
