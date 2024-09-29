package graphs

import "bulentkeskin.com/dsa/library"

func HasPath(graph Graph[rune], src rune, dst rune) bool {
	var stack library.Stack[rune]
	stack.Push(src)

	for stack.Len() > 0 {
		curr := stack.Pop()

		if curr == dst {
			return true
		}

		for _, child := range graph[curr] {
			stack.Push(child)
		}
	}

	return false
}
