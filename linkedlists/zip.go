package linkedlists

/*
Write a function that takes in the head of two linked lists as arguments.
The function should zipper the two lists together into single linked list by alternating nodes.
If one of the linked lists is longer than the other, the resulting list should terminate with the remaining nodes.
The function should return the head of the zippered linked list.

Do this in-place, by mutating the original Nodes.
*/

func Zip(list1 *Node, list2 *Node) *Node {
	head := list1
	if head == nil {
		head = list2
	}

	for list1 != nil && list2 != nil {
		list1Next := list1.Next
		list2Next := list2.Next
		list1.Next = list2
		if list1Next != nil {
			list2.Next = list1Next
		}
		list2 = list2Next
		list1 = list1Next
	}

	return head
}
