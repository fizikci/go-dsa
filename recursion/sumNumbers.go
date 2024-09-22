package recursion

func SumNumbers(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + SumNumbers(arr[1:])
}
