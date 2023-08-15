package main

/*
	LC #189
	TC---> O(n)
	SC--->O(1)
*/
func rotate(arr []int, k int) {
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

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
