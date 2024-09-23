package linkedlists

/*
Write a function that takes in the head of a linked list as an argument.
The function should reverse the order of the nodes in the linked list in-place and return the new head of the reversed linked list.
*/

func Reverse(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}

	return prev
}
