package main

func findSubstring(s string, words []string) []int {
	stringLen := len(s)
	wordArrayLen := len(words)
	wordLen := len(words[0])
	substringLen := wordArrayLen * wordLen
	wordCount := make(map[string]int)
	res := []int{}

	for _, word := range words {
		wordCount[word] += 1
	}

	slidingWindow := func(left int) {
		excessWord := false
		wordsFound := make(map[string]int)
		wordsUsed := 0

		for right := left; right <= stringLen-wordLen; right += wordLen {
			subStr := s[right : right+wordLen]

			if count, exist := wordCount[subStr]; !exist {
				// reset everything
				excessWord = false
				wordsUsed = 0
				wordsFound = make(map[string]int)
				left = right + wordLen
			} else {
				for right-left == substringLen || excessWord {
					leftMostWord := s[left : left+wordLen]
					wordsFound[leftMostWord] -= 1
					left = left + wordLen

					if wordsFound[leftMostWord] == wordCount[leftMostWord] {
						excessWord = false
					} else {
						wordsUsed -= 1
					}
				}

				wordsFound[subStr] += 1

				if wordsFound[subStr] <= count {
					wordsUsed++
				} else {
					excessWord = true
				}

				if wordsUsed == wordArrayLen && !excessWord {
					res = append(res, left)
				}

			}

		}
	}

	for i := 0; i < wordLen; i++ {
		slidingWindow(i)
	}

	return res
}
