package linkedlists

/*
Write a function that takes in the head of a linked list, a value, and an index.
The function should insert a new node with the value into the list at the specified index.
Consider the head of the linked list as index 0. The function should return the head of the resulting linked list.

Do this in-place.
*/

func InsertNode(head *Node, val int, index int) *Node {
	if index == 0 {
		return &Node{val, head}
	}

	curr := head

	for i := 0; i < index-1; i++ {
		curr = curr.Next
		if curr == nil {
			break
		}
	}

	if curr != nil {
		curr.Next = &Node{val, curr.Next}
	}

	return head
}
