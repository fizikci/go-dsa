package linkedlists

/*
Write a function that takes in the head of a linked list and a target value as arguments.
The function should delete the node containing the target value from the linked list and
return the head of the resulting linked list. If the target appears multiple times in the linked list,
only remove the first instance of the target in the list.

Do this in-place.
*/

func RemoveNode(head *Node, val int) *Node {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}

	curr := head
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}

	return head
}
