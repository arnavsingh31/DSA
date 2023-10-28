package twopointer

func TwoSum(arr []int, target int) []int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		sum := arr[start] + arr[end]
		if sum == target {
			break
		} else if sum > target {
			end--
		} else if sum < target {
			start++
		}
	}
	return []int{start + 1, end + 1}
}
