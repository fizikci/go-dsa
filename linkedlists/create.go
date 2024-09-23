package linkedlists

/*
Write a function, createLinkedList, that takes in an array of values as an argument.
The function should create a linked list containing each element of the array as the values of the nodes.
The function should return the head of the linked list.
*/

func Create(slice []int) *Node {
	dummy := &Node{-1, nil}
	curr := dummy

	for _, v := range slice {
		curr.Next = &Node{v, nil}
		curr = curr.Next
	}

	return dummy.Next
}
