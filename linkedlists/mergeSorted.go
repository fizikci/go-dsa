package linkedlists

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
