package arrays

/*
Write a function that takes in two arrays, a,b, as arguments.
The function should return a new array containing elements that are in both of the two arrays.
*/

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
