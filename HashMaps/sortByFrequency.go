package hashmaps

import (
	"sort"
)

/*
LC #451
TC--->O(nlogn)
SC--->O(n)
Brute force {Accepted}
*/
func SortCharcterByFrequency(s string) string {
	freqMap := make(map[byte]int)
	arr := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		freqMap[s[i]] += 1
		arr = append(arr, s[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		if freqMap[arr[i]] == freqMap[arr[j]] {
			return arr[i] > arr[j]
		}

		return freqMap[arr[i]] > freqMap[arr[j]]
	})

	return string(arr)
}

/*
TC--->O(n + klogk), k is the number of unique characters in string
SC--->O(n)
*/
func SortCharcterByFrequency2(s string) string {
	freqMap := make(map[byte]int)
	heap := make(MaxHeap, 0)
	ans := ""

	for i := 0; i < len(s); i++ {
		freqMap[s[i]] += 1
	}

	for key, value := range freqMap {
		heap.Push(CharInfo{
			char:      key,
			frequency: value,
		})
	}

	for len(heap) > 0 {
		ci := heap.top()
		heap.Pop()

		for i := 0; i < ci.frequency; i++ {
			ans += string(ci.char)
		}
	}

	return ans
}

type MaxHeap []CharInfo
type CharInfo struct {
	char      byte
	frequency int
}

func (mh *MaxHeap) Push(ci CharInfo) {
	*mh = append(*mh, ci)

	index := len(*mh) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if (*mh)[parent].frequency < (*mh)[index].frequency {
			mh.Swap(index, parent)
			index = parent
		} else {
			return
		}
	}
}

func (mh *MaxHeap) Pop() {
	lastIndex := len(*mh) - 1

	mh.Swap(0, lastIndex)
	(*mh) = (*mh)[:lastIndex]

	// call heapify
	mh.Heapify(0, len(*mh)-1)
}

func (mh *MaxHeap) Heapify(index, n int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left <= n && (*mh)[largest].frequency < (*mh)[left].frequency {
		largest = left
	}

	if left <= n && (*mh)[largest].frequency == (*mh)[left].frequency {
		if (*mh)[largest].char < (*mh)[left].char {
			largest = left
		}
	}

	if right <= n && (*mh)[largest].frequency < (*mh)[right].frequency {
		largest = right
	}

	if right <= n && (*mh)[largest].frequency == (*mh)[right].frequency {
		if (*mh)[largest].char < (*mh)[right].char {
			largest = right
		}
	}

	if largest != index {
		mh.Swap(largest, index)
		mh.Heapify(largest, n)
	}
}

func (mh *MaxHeap) top() CharInfo {
	return (*mh)[0]
}

func (mh *MaxHeap) Swap(i, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}
