package library

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Slice[T any] []T

func (arr Slice[T]) Find(predicate func(T) bool) (T, bool) {
	for _, v := range arr {
		if predicate(v) {
			return v, true
		}
	}

	var zeroValue T
	return zeroValue, false
}
