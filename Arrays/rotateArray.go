package arrays

/*
	LC #189
	TC---> O(n)
	SC--->O(1)
*/
func Rotate(arr []int, k int) {
	k = k % len(arr)

	reverse(&arr, 0, len(arr)-1)
	reverse(&arr, 0, k-1)
	reverse(&arr, k, len(arr)-1)

}

func reverse(arr *[]int, start, end int) {
	for start < end {
		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
		start++
		end--
	}
}
