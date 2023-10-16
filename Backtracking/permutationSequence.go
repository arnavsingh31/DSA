package backtracking

import "fmt"

/*
	LC #60
	TC--->O(n * n)
	SC--->O(n)
*/
func GetKthPermutation(n, k int) string {
	arr := make([]int, 0)
	fact := 1
	for i := 1; i < n; i++ {
		fact = fact * i
		arr = append(arr, i)
	}
	arr = append(arr, n)
	k = k - 1
	ans := ""

	for {
		index := k / fact
		ans += fmt.Sprint(arr[index])
		removeFromArray(&arr, index)

		if len(arr) == 0 {
			break
		}

		k = k % fact
		fact = fact / len(arr)
	}

	return ans
}

func removeFromArray(arr *[]int, index int) {
	for i := index; i < len(*arr)-1; i++ {
		(*arr)[i] = (*arr)[i+1]
	}

	*arr = (*arr)[:len(*arr)-1]
}
