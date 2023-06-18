package main

import (
	"strings"
)

type IntRoman struct {
	IntVal   int
	RomanVal string
}

func intToRoman(num int) string {
	arr := []IntRoman{{
		IntVal:   1000,
		RomanVal: "M",
	}, {
		IntVal:   900,
		RomanVal: "CM",
	}, {
		IntVal:   500,
		RomanVal: "D",
	}, {
		IntVal:   400,
		RomanVal: "CD",
	}, {
		IntVal:   100,
		RomanVal: "C",
	}, {
		IntVal:   90,
		RomanVal: "XC",
	}, {
		IntVal:   50,
		RomanVal: "L",
	}, {
		IntVal:   40,
		RomanVal: "XL",
	}, {
		IntVal:   10,
		RomanVal: "X",
	}, {
		IntVal:   9,
		RomanVal: "IX",
	}, {
		IntVal:   5,
		RomanVal: "V",
	}, {
		IntVal:   4,
		RomanVal: "IV",
	}, {
		IntVal:   1,
		RomanVal: "I",
	}}

	romanStr := ""

	for i := 0; i < len(arr); i++ {
		if num >= arr[i].IntVal {
			counting := num / arr[i].IntVal
			romanStr += strings.Repeat(arr[i].RomanVal, counting)
			num %= arr[i].IntVal
		}
	}

	return romanStr
}
