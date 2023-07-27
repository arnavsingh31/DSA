package main

/*
	LC #1574
	Explanation: https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/solutions/830416/java-increasing-from-left-right-and-merge-o-n/
*/
func findLengthOfSubArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < len(arr)-1 && arr[left] <= arr[left+1] {
		left++
	}
	if left == len(arr)-1 {
		return 0
	}

	for right > left && arr[right] >= arr[right-1] {
		right--
	}
	res := minLen(right, len(arr)-left-1)

	i := 0
	j := right
	for i <= left && j <= len(arr)-1 {
		if arr[j] >= arr[i] {
			res = minLen(res, j-i-1)
			i++
		} else {
			j++
		}
	}

	return res
}

func minLen(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	log.Print(findLengthOfSubArray([]int{5, 4, 3, 2, 1}))
// }
