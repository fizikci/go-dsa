package arrays

import "strconv"

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
