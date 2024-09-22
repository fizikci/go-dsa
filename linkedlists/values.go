package linkedlists

func Values(head *Node) []int {
	res := make([]int, 0, 5)

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
