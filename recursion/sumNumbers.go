package recursion

/*
Write a function sumOfLengths that takes in array of strings and returns the total length of the strings.
*/

func SumNumbers(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + SumNumbers(arr[1:])
}
