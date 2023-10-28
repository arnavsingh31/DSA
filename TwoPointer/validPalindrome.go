package twopointer

import (
	"regexp"
	"strings"
)

func ValidPalindrome(s string) bool {
	s = strings.ToLower(s)
	newStr := ""

	for _, char := range s {
		if isValid(byte(char)) {
			newStr += string(char)
		}
	}

	if newStr == "" {
		return true
	}

	start := 0
	end := len(newStr) - 1

	for start <= end {
		if newStr[start] != newStr[end] {
			return false
		}
		start++
		end--
	}
	return true

}

// this code is shorter version of the code above.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		for left < right && !isValid(s[left]) {
			left++
		}
		for left < right && !isValid(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func isValid(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

// using regexp
func ValidPalindrome2(s string) bool {
	s = strings.ToLower(s)
	newStr := clearString(s)

	if newStr == "" {
		return true
	}

	start := 0
	end := len(newStr) - 1

	for start <= end {
		if newStr[start] != newStr[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func clearString(s string) string {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

	return nonAlphanumericRegex.ReplaceAllString(s, "")

}
