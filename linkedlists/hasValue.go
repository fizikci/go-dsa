package linkedlists

func HasValue(head *Node, val int) bool {
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == val {
			return true
		}
	}

	return false
}
