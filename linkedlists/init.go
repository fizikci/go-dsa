package linkedlists

import "strconv"

type Node struct {
	Val  int
	Next *Node
}

func (head *Node) String() (res string) {
	for curr := head; curr != nil; curr = curr.Next {
		res += strconv.Itoa(curr.Val)
		if curr.Next != nil {
			res += " -> "
		}
	}

	return res
}

var (
	IntLinkedList  *Node
	IntLinkedList2 *Node
)

func init() {
	IntLinkedList = &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	IntLinkedList2 = &Node{100, &Node{99, nil}}
}
