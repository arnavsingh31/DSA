package main

/*
	LC #229
	TC--->O(n)
	SC--->O(n)
	Using hashmap
*/
func majorityElement(nums []int) []int {
	numCountMap := make(map[int]int)
	n := len(nums)
	ans := []int{}

	for i := 0; i < n; i++ {
		numCountMap[nums[i]] += 1
	}

	for key, value := range numCountMap {
		if value > n/3 {
			ans = append(ans, key)
		}
	}
	return ans
}

// Todo: Implement using boyer-moore voting algorithm
