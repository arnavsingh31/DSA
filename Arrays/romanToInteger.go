package arrays

func RomanToInt(s string) int {
	roman_To_Int_Map := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	integer := 0

	for i := 0; i < len(s); i++ {
		curr := roman_To_Int_Map[s[i]]
		if i+1 < len(s) && curr < roman_To_Int_Map[s[i+1]] {
			integer -= curr
		} else {
			integer += curr
		}
	}

	return integer
}
