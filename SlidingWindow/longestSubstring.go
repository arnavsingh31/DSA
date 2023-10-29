package slidingwindow

import util "github.com/arnavsingh31/DSA/Util"

func LongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen := 1
	left := 0

	for right := 1; right < len(s); right++ {
		curr := s[right]

		for i := left; i < right; i++ {
			if curr == s[i] {
				left = i + 1
				break
			}
		}

		maxLen = util.Max(maxLen, right-left+1)
	}

	return maxLen
}

// Optimsed
func LongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left := 0
	maxLen := 1
	hashMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		curr := s[right]

		if prevIndex, exist := hashMap[curr]; exist {

			if prevIndex >= left {
				left = prevIndex + 1
			}
		}
		hashMap[curr] = right

		maxLen = util.Max(maxLen, right-left+1)
	}

	return maxLen

}
