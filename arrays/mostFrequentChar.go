package arrays

/*
Write a function that takes a string and returns the character that appears the most frequently.
*/

func MostFrequentChar(s string) rune {
	count := make(map[rune]int)

	max := 0
	var char rune

	for _, c := range s {
		count[c]++

		if count[c] > max {
			max = count[c]
			char = c
		}
	}

	return char
}
