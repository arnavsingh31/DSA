package binarysearch

/*
	LC #1539
	TC--->O(n)
	SC--->O(n)
*/
func KthMissingPositiveNumber(nums []int, k int) int {
	numMap := make(map[int]bool)
	missingNum := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true
	}

	for counter := 1; len(missingNum) < k; counter++ {
		if !numMap[counter] {
			missingNum = append(missingNum, counter)
		}
	}

	return missingNum[k-1]
}

/*
	TC--->O(n)
	SC--->O(1)
	Better solution than above
*/
func KthMissingPositiveNumber2(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= k {
			k++
		} else {
			break
		}
	}

	return k
}

/*
	TODO--Optimise using Binary Search
	TC--->O()
	SC--->O()
*/
// func KthMissingPositiveNumber3(nums []int, k int) int {
// }
