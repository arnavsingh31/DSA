package backtracking

/*
	TC--->O(n*n)
	SC--->O(1)
*/
func InsertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}
	}

	return arr
}
