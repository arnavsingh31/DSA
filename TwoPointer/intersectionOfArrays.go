package twopointer

import "sort"

// T.C---> O(nlogn) S.C---> O(1) if both arrays are sorted then T.C--> O(n)
func IntersectionOfArrays(nums1, nums2 []int) []int {
	ans := []int{}

	sort.Ints(nums1)
	sort.Ints(nums2)

	ptr1, ptr2 := 0, 0

	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] == nums2[ptr2] {
			if len(ans) == 0 || nums1[ptr1] != ans[len(ans)-1] {
				ans = append(ans, nums1[ptr1])
			}
			ptr1++
			ptr2++
		} else if nums1[ptr1] > nums2[ptr2] {
			ptr2++
		} else if nums1[ptr1] < nums2[ptr2] {
			ptr1++
		}
	}
	return ans
}

// T.C---> O(n)  S.C--->O(n)
func IntersectionOfArrays2(nums1, nums2 []int) []int {
	ans := []int{}

	array_map := make(map[int]bool)

	for _, num := range nums1 {
		array_map[num] = true
	}

	for _, num := range nums2 {
		if array_map[num] {
			ans = append(ans, num)
			array_map[num] = false
		}
	}
	return ans
}
