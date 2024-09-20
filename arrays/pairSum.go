package arrays

/*
Write a function, PairSum, that takes in an array of numbers and a target number. The function should return
a pair of indices whose elements sum to the given target. The indices should be ordered in ascending order.
*/

func PairSum(arr []int, target int) []int {
	numIndexes := make(map[int]int)

	for i, num := range arr {
		complement := target - num

		if index, ok := numIndexes[complement]; ok {
			return []int{index, i}
		}

		numIndexes[num] = i
	}

	return []int{}
}
