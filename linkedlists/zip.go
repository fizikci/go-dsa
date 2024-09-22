package linkedlists

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
