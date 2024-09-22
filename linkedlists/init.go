package linkedlists

type Node struct {
	Val  int
	Next *Node
}

var (
	IntLinkedList *Node
)

func init() {
	IntLinkedList = &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
}
