package linkedlists

/*
Write a function that takes in the head of two sorted linked lists as arguments.
The function should merge the two lists together into single sorted linked list.
The function should return the head of the merged linked list.

Do this in-place, by mutating the original Nodes.
*/

func MergeSorted(list1 *Node, list2 *Node) *Node {
	var dummy = &Node{-1, nil}

	curr := dummy

	for curr != nil {
		if list1 == nil {
			curr.Next = list2
			break
		}
		if list2 == nil {
			curr.Next = list1
			break
		}

		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	return dummy.Next
}
