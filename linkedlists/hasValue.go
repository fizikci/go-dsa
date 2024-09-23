package linkedlists

/*
Write a function that takes in the head of a linked list and a target value.
The function should return a boolean indicating whether or not the linked list contains the target.
*/

func HasValue(head *Node, val int) bool {
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			return true
		}
	}

	return false
}
