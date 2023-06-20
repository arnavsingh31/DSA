package main

func reverseWords(s string) string {
	i := len(s) - 1
	var j int
	ans := ""

	for i >= 0 {
		for i >= 0 && s[i] == 32 {
			i--
		}
		j = i
		// to avoid concatenating extra space in case where we have leading spaces.
		if i < 0 {
			break
		}

		for i >= 0 && s[i] != 32 {
			i--
		}
		if len(ans) == 0 {
			ans += s[i+1 : j+1]
		} else {
			ans += " " + s[i+1:j+1]
		}

	}
	return ans
}
