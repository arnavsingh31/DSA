package backtracking

/*
	TC--->O(nlogn)
	SC--->O(n)
*/
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	ptr1, ptr2 := 0, 0
	mergeArr := make([]int, 0)

	for ptr1 < len(arr1) && ptr2 < len(arr2) {
		if arr1[ptr1] < arr2[ptr2] {
			mergeArr = append(mergeArr, arr1[ptr1])
			ptr1++
		} else if arr1[ptr1] > arr2[ptr2] {
			mergeArr = append(mergeArr, arr2[ptr2])
			ptr2++
		} else {
			mergeArr = append(mergeArr, arr1[ptr1])
			mergeArr = append(mergeArr, arr2[ptr2])
			ptr1++
			ptr2++
		}
	}

	for ptr1 < len(arr1) {
		mergeArr = append(mergeArr, arr1[ptr1])
		ptr1++
	}

	for ptr2 < len(arr2) {
		mergeArr = append(mergeArr, arr2[ptr2])
		ptr2++
	}

	return mergeArr
}
