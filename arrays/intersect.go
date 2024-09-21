package arrays

func Intersect(arr1, arr2 []int) []int {
	m := make(map[int]bool)

	for _, v := range arr1 {
		m[v] = true
	}

	res := []int{}

	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}

	return res
}
