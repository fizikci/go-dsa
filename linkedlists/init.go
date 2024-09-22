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
	IntLinkedList = &Node{10, &Node{20, &Node{30, &Node{40, &Node{40, &Node{40, &Node{50, nil}}}}}}}
	IntLinkedList2 = &Node{24, &Node{25, nil}}
}
