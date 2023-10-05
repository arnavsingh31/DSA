package maths

/*
	LC #78
	TC--->(2^n * n)
	SC--->O(1)
*/
func Subsets(arr []int) [][]int {
	ans := make([][]int, 0)
	n := len(arr)

	for i := 0; i < 1<<n; i++ {
		temp := make([]int, 0)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				temp = append(temp, arr[j])
			}
		}
		ans = append(ans, temp)
	}

	return ans
}

func Subsets2(nums []int) [][]int {
	ans := make([][]int, 0)

	solveRec(nums, 0, &ans, []int{})
	return ans
}

func solveRec(nums []int, index int, ans *[][]int, arr []int) {
	if index >= len(nums) {
		temp := append([]int{}, arr...)
		*ans = append(*ans, temp)
		return
	}

	// exclude
	solveRec(nums, index+1, ans, arr)

	// include
	arr = append(arr, nums[index])
	solveRec(nums, index+1, ans, arr)
}
