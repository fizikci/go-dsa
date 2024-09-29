package graphs

import (
	"slices"

	"bulentkeskin.com/dsa/library"
)

func ConnectedComponentCount(graph Graph[int]) int {

	getComponent := func(key int) []int {
		var component []int
		var stack library.Stack[int]
		stack.Push(key)

		for stack.Len() > 0 {
			curr := stack.Pop()
			if slices.Contains(component, curr) {
				continue
			}

			component = append(component, curr)

			for _, k := range graph[curr] {
				stack.Push(k)
			}
		}

		return component
	}

	var components library.Slice[[]int]

	for key := range graph {
		if _, found := components.Find(func(comp []int) bool { return slices.Contains(comp, key) }); found {
			continue
		}

		components = append(components, getComponent(key))
	}

	return len(components)
}
