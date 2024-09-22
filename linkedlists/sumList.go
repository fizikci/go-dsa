package linkedlists

func SumList(head *Node) int {
	sum := 0

	for curr := head; curr != nil; curr = curr.Next {
		sum += curr.Val
	}

	return sum
}
