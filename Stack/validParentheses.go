package main

// LC #20 T.C--->O(n), S.C--->O(n)
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	parenthesesMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, p := range s {
		if _, exist := parenthesesMap[p]; exist {
			stack = append(stack, p)

		} else if len(stack) == 0 || parenthesesMap[stack[len(stack)-1]] != p {
			return false
		} else {
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}
