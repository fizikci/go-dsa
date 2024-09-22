package linkedlists

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
