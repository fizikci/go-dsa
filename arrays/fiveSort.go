package arrays

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
