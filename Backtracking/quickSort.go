package backtracking

/*
	TC--->O(nlogn)
	SC--->O(1)
*/
func QuickSort(arr []int) []int {
	qs(arr, 0, len(arr)-1)
	return arr
}

func qs(arr []int, low, high int) {
	// if low == high then we have only one element implying it is already sorted
	if low < high {
		partionIndex := sortRec(&arr, low, high)
		qs(arr, low, partionIndex-1)
		qs(arr, partionIndex+1, high)
	}
}

func sortRec(arr *[]int, low, high int) int {
	// we take our first element in the range as our pivot
	pivot := (*arr)[low]
	i := low
	j := high
	for i < j {
		// below loops finds index of first element > pivot
		for i < high && pivot >= (*arr)[i] {
			i++
		}

		// below loops finds index of first element < pivot
		for j > low && pivot < (*arr)[j] {
			j--
		}

		// swap only if i < j, if i has crossed j then we have found the correct position for our pivot which is j
		if i < j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	// swap arr[j] with pivot which is at low
	(*arr)[j], (*arr)[low] = (*arr)[low], (*arr)[j]

	// j is the partition index
	return j
}
