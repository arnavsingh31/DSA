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

/*
	TC--->O(n)
	SC--->O(1)
	Using Boyer-Moore Majority voting algorithm
*/
func majorityElement2(nums []int) []int {
	ans := []int{}
	n := len(nums)
	candidate1 := -1
	candidate2 := -1
	count1 := 0
	count2 := 0

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}

		if count1 == 0 {
			candidate1 = num
			count1++
		} else if count2 == 0 {
			candidate2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	// now check the candidature of both canditates
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	if count1 > n/3 {
		ans = append(ans, candidate1)
	}
	if count2 > n/3 {
		ans = append(ans, candidate2)
	}

	return ans
}
