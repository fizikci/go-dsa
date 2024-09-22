package linkedlists

/*
Write a function, longestStreak, that takes in the head of a linked list as an argument.
The function should return the length of the longest consecutive streak of the same value within the list.
*/

func LongestStreak(head *Node) int {
	longest := 0
	currLen := 1

	for curr := head; curr != nil && curr.Next != nil; curr = curr.Next {
		if curr.Val == curr.Next.Val {
			currLen++
			if longest < currLen {
				longest = currLen
			}
		} else {
			currLen = 1
		}
	}

	return longest
}
