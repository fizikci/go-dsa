package arrays

import (
	"strconv"
	"strings"
	"unicode"
)

/*
Write a function that takes in a string as an argument. The input string will be formatted into
multiple groups according to the following rules:
*/

func Uncompress(s string) string {
	var result string
	var lastNum string

	for _, c := range s {
		if unicode.IsDigit(c) {
			lastNum += string(c)
		} else {
			num, _ := strconv.Atoi(lastNum)
			result += strings.Repeat(string(c), num)
			lastNum = ""
		}
	}

	return result
}
