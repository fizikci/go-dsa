package arrays

/*
Write a function that takes in an array of numbers as an argument.
The function should rearrange elements of the array such that all 5s appear at the end.
Your function should perform this operation in-place by mutating the original array. The function should return the array.

Elements that are not 5 can appear in any order in the output, as long as all 5s are at the end of the array.
*/

func FiveSort(arr []int) []int {
	i, j := 0, len(arr)-1

	for i < j {
		if arr[i] != 5 {
			i++
			continue
		}
		if arr[j] == 5 {
			j--
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
