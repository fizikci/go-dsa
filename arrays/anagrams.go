package arrays

/*
Write a function that takes in two strings as arguments. The function should return a boolean
indicating whether or not the strings are anagrams. Anagrams are strings that contain the same characters,
but in any order.
*/

func Anagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)

	for _, r := range s1 {
		counts[r]++
	}

	for _, r := range s2 {
		counts[r]--
		if counts[r] < 0 {
			return false
		}
	}

	return true
}
