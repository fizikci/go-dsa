package linkedlists

func Reverse(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
	}

	return prev
}
