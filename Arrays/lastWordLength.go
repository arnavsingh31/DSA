package arrays

func LastWordLen(s string) int {
	start, end := -1, -1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 && end == -1 {
			end = i
		}

		if s[i] == 32 && end != -1 {
			start = i + 1
			break
		}

		if s[i] != 32 && i == 0 {
			start = i
		}
	}

	return end - start + 1
}
