package twopointer

func IsSubsequence(s, t string) bool {
	if len(t) < len(s) {
		return false
	}

	sPtr := 0
	tPtr := 0
	temp := ""

	for sPtr < len(s) && tPtr < len(t) {
		if s[sPtr] == t[tPtr] {
			temp += string(s[sPtr])
			sPtr++
		}
		tPtr++
	}

	return len(s) == len(temp)
}
