package arrays

import "strconv"

/*
Write a function that takes in a string as an argument. The function should return
a compressed version of the string where consecutive occurrences of the same characters are
compressed into the number of occurrences followed by the character. Single character occurrences should not be changed.

'aaa' compresses to '3a'
'cc' compresses to '2c'
't' should remain as 't'
*/

func Compress(s string) string {
	var result string
	arr := []rune(s)
	lastRune := arr[0]
	count := 1

	for i := 1; i < len(arr); i++ {
		if lastRune == arr[i] {
			count++
		} else {
			result += createCompressedToken(count, lastRune)
			lastRune = arr[i]
			count = 1
		}
	}

	result += createCompressedToken(count, lastRune)

	return result
}

func createCompressedToken(count int, r rune) string {
	if count == 1 {
		return string(r)
	} else {
		return strconv.Itoa(count) + string(r)
	}
}
