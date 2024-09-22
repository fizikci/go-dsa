package linkedlists

func Create(slice []int) *Node {
	dummy := &Node{-1, nil}
	curr := dummy

	for _, v := range slice {
		curr.Next = &Node{v, nil}
		curr = curr.Next
	}

	return dummy.Next
}
