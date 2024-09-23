package recursion

/*
Write a function that takes in a string as an argument.
The function should return the string with its characters in reverse order.
*/

func ReverseString(s string) string {
	if s == "" {
		return ""
	}

	return ReverseString(s[1:]) + string(s[0])
}
