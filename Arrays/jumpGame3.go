package main

/*
	LC #1306
	TC--->O(n)
	SC--->O(n)
*/
func jumpGame3(arr []int, start int) bool {
	queue := []int{start}

	for len(queue) > 0 {
		pos, val := queue[0], arr[queue[0]]
		queue = queue[1:]

		if val == 0 {
			return true
		}

		// if val is negative we know we have already visited this index, so skip it.
		if val < 0 {
			continue
		}

		if pos-arr[pos] >= 0 {
			queue = append(queue, pos-arr[pos])
		}

		if pos+arr[pos] < len(arr) {
			queue = append(queue, pos+arr[pos])
		}

		// multiplying val by -1 in- order to skip it if same index appears multiple times.
		arr[pos] *= -1
	}
	return false
}
