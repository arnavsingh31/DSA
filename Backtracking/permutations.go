package backtracking

/*
	LC #
	TC--->O(n! * n)
	SC--->O(n{curr permutation} + n {hashmap})
*/
func Permute(nums []int) [][]int {
	isTaken := make(map[int]bool)
	ans := make([][]int, 0)
	solveRecPer(nums, []int{}, isTaken, &ans)
	return ans
}

func solveRecPer(arr []int, per []int, isTaken map[int]bool, ans *[][]int) {
	if len(per) == len(arr) {
		temp := append([]int{}, per...)
		*ans = append(*ans, temp)
		return
	}

	for i := 0; i < len(arr); i++ {
		if _, exist := isTaken[i]; !exist {
			isTaken[i] = true
			per = append(per, arr[i])
			solveRecPer(arr, per, isTaken, ans)
			delete(isTaken, i)
			per = per[:len(per)-1]
		}
	}
}

func Permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	solveRecPer2(nums, 0, &ans)
	return ans
}

func solveRecPer2(arr []int, index int, ans *[][]int) {
	if index >= len(arr) {
		temp := append([]int{}, arr...)
		*ans = append(*ans, temp)
		return
	}

	for i := index; i < len(arr); i++ {
		arr[i], arr[index] = arr[index], arr[i]
		solveRecPer2(arr, index+1, ans)
		arr[i], arr[index] = arr[index], arr[i]
	}
}
