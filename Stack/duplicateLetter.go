package stack

/*
	LC #316
	TC--->O(n)
	SC--->O(n)
	Alternate way of asking question: LC #1081 Smallest Subsequence of Distinct Characters
*/
func RemoveDuplicateLetter(s string) string {
	stack := []rune{}
	freqMap := make(map[rune]int)
	stackContainsMap := make(map[rune]bool)

	// populate frequency map
	for _, r := range s {
		freqMap[r] += 1
	}

	for _, r := range s {
		_, contains := stackContainsMap[r]

		for !contains && len(stack) > 0 && stack[len(stack)-1] >= r && freqMap[stack[len(stack)-1]] > 0 {
			delete(stackContainsMap, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		if !contains {
			stack = append(stack, r)
			stackContainsMap[r] = true
		}

		freqMap[r] -= 1
	}

	return string(stack)

}
