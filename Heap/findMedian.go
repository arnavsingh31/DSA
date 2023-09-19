package heap

import util "github.com/arnavsingh31/DSA/Util"

/*
	VERY-HARD to come up with solution. For future reference watch video.
	TC--->O(nlogn)
	SC--->O(n)
*/
func FindMedian(arr []int) []float64 {
	minHeap := make(MinHeap, 0)
	maxHeap := make(MaxHeap, 0)
	median := 0.0
	ans := []float64{}

	for i := 0; i < len(arr); i++ {
		calculateMedian(arr[i], &minHeap, &maxHeap, &median)
		ans = append(ans, median)
	}

	return ans
}

func calculateMedian(element int, minHeap *MinHeap, maxHeap *MaxHeap, median *float64) {
	switch util.Signum(len(*maxHeap), len(*minHeap)) {
	case 0:
		if float64(element) < *median {
			maxHeap.push(element)
			*median = float64((*maxHeap).top())
		} else {
			minHeap.push(element)
			*median = float64((*minHeap).top())
		}

	case 1:
		if float64(element) > *median {
			minHeap.push(element)
			*median = util.Average((*minHeap).top(), (*maxHeap).top())
		} else {
			minHeap.push((*maxHeap).top())
			maxHeap.pop()
			maxHeap.push(element)
			*median = util.Average((*minHeap).top(), (*maxHeap).top())
		}

	case -1:
		if float64(element) > *median {
			maxHeap.push((*minHeap).top())
			minHeap.pop()
			minHeap.push(element)
			*median = util.Average((*minHeap).top(), (*maxHeap).top())
		} else {
			maxHeap.push(element)
			*median = util.Average((*minHeap).top(), (*maxHeap).top())
		}

	}
}
