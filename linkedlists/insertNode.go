package linkedlists

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
