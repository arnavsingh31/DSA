package main

/*
	NOTE:
	Implemented heap following 1 based indexing.
	If you want to follow zero based indexing then formula for left and right child of a node will
	change left := 2*i + 1 , right := 2*i + 2.
*/

func insertInHeap(arr *[]int, val int) {
	// first insert at end
	*arr = append(*arr, val)
	index := len(*arr) - 1

	//then find correct position of the val {here using max heap to find val position}
	for index > 1 {
		parent := (index) / 2
		if (*arr)[index] > (*arr)[parent] {
			(*arr)[index], (*arr)[parent] = (*arr)[parent], (*arr)[index]
			index = parent
		} else {
			return
		}
	}
}

// we always delete a root node.
func deleteFromHeap(arr *[]int) {

	// first copy last node to first node
	(*arr)[1] = (*arr)[len(*arr)-1]

	// delete last node
	*arr = (*arr)[:len(*arr)-1]
	n := len(*arr)
	// now find the correct position of root in heap
	index := 1
	for index < n {
		left := 2 * index
		right := 2*index + 1
		swapIndex := index

		if left < n && (*arr)[swapIndex] < (*arr)[left] {
			swapIndex = left
		}

		if right < n && (*arr)[swapIndex] < (*arr)[right] {
			swapIndex = right
		}

		if swapIndex != index {
			(*arr)[index], (*arr)[swapIndex] = (*arr)[swapIndex], (*arr)[index]
		} else {
			return
		}

	}
}

// heapify algorithm {convert given array into max/min heap}. Below we implemented max heap.
func heapify(arr *[]int, i int) {
	left := 2 * i
	right := 2*i + 1
	swapIndex := i

	if left < len(*arr) && (*arr)[swapIndex] < (*arr)[left] {
		swapIndex = left
	}

	if right < len(*arr) && (*arr)[swapIndex] < (*arr)[right] {
		swapIndex = right
	}

	if swapIndex != i {
		(*arr)[i], (*arr)[swapIndex] = (*arr)[swapIndex], (*arr)[i]
		heapify(arr, swapIndex)
	}

}

// min heap
func heapify2(arr *[]int, i int) {
	left := 2 * i
	right := 2*i + 1
	swapIndex := i

	if left < len(*arr) && (*arr)[swapIndex] > (*arr)[left] {
		swapIndex = left
	}

	if right < len(*arr) && (*arr)[swapIndex] > (*arr)[right] {
		swapIndex = right
	}

	if swapIndex != i {
		(*arr)[i], (*arr)[swapIndex] = (*arr)[swapIndex], (*arr)[i]
		heapify2(arr, swapIndex)
	}

}

// func main() {
// arr := make([]int, 1)
// arr[0] = -1
// insertInHeap(&arr, 54)
// insertInHeap(&arr, 53)
// insertInHeap(&arr, 55)
// insertInHeap(&arr, 52)
// insertInHeap(&arr, 50)
// log.Print(arr)

// deleteFromHeap(&arr)
// log.Print(arr)

// heapify algorithm TC--->O(n) how?
// arr2 := []int{-1, 54, 53, 55, 52, 50}
// n := len(arr2)
// for i := (n - 1) / 2; i > 0; i-- {
// 	heapify2(&arr2, i)
// }

// log.Print(arr2)
// }

/*
	 50
	/  \
   53   55
  /  \
 52  54


*/
