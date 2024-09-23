package linkedlists

/*
Write a function that takes in the head of a linked list as an argument.
The function should return an array containing all values of the nodes in the linked list.
*/

func Values(head *Node) []int {
	res := make([]int, 0, 5)

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
