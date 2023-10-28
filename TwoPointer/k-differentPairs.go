package twopointer

import (
	"math"
	"sort"
)

/*
	T.C---> O(n) S.C--> O(n)
*/

func KDifferentPairs(nums []int, k int) int {
	numFreqMap := make(map[int]int)

	for _, num := range nums {
		numFreqMap[num] += 1
	}
	ans := 0
	if k == 0 {
		for _, val := range numFreqMap {
			if val > 0 {
				ans++
			}
		}
	} else {
		for key, _ := range numFreqMap {
			if _, exist := numFreqMap[key+k]; exist {
				ans++
			}
		}
	}
	return ans
}

/*
T.C--> O(nlogn) S.C---> O(1)
https://leetcode.com/problems/k-diff-pairs-in-an-array/solutions/1756874/c-multiple-approaches-maps-two-pointer/
*/
func KDifferentPairs2(nums []int, k int) int {
	sort.Ints(nums)

	left, right := 0, 1
	count := 0

	for left < len(nums) && right < len(nums) {
		diff := int(math.Abs(float64(nums[left] - nums[right])))
		if left == right || diff < k {
			right++
		} else {
			if diff == k {
				count++
				right++
				for right < len(nums) && nums[right] == nums[right-1] {
					right++
				}
			}
			left++
			for left < len(nums) && nums[left] == nums[left-1] {
				left++
			}
		}
	}
	return count
}
