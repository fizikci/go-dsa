package linkedlists

/*
Write a function that takes in the head of a linked list containing numbers as an argument.
The function should return the total sum of all values in the linked list.
*/

func SumList(head *Node) int {
	sum := 0

	for curr := head; curr != nil; curr = curr.Next {
		sum += curr.Val
	}

	return sum
}
