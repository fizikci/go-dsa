package recursion

/*
Write a function that takes in a string and returns a boolean indicating whether or not the string is the same forwards and backwards.
*/

func IsPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return IsPalindrome(s[1 : len(s)-1])
}
