package main

import (
	"strings"
)

func wordPattern(pattern, s string) bool {
	words := strings.Fields(s)

	if len(words) != len(pattern) {
		return false
	}

	mapCharToWord := make(map[string]string)
	usedWordMap := make(map[string]bool)

	for i := 0; i < len(pattern); i++ {
		if mappedWord, exist := mapCharToWord[string(pattern[i])]; exist {
			if mappedWord != words[i] {
				return false
			}
		} else {
			if usedWordMap[words[i]] {
				return false
			}
			mapCharToWord[string(pattern[i])] = words[i]
			usedWordMap[words[i]] = true
		}
	}
	return true
}
