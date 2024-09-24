package twopointer

import "sort"

// LC #18
// TC :- O(n^3)
// SC :- O(1) no extra space used apart from the ans
func FourSum(nums []int, target int) [][]int {
	n := len(nums) - 1
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := n - 1

			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]

				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--

					for k < l && nums[k] == nums[k-1] {
						k++
					}

					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum > target {
					l--
				} else {
					k++
				}
			}
		}
	}
	return ans
}
