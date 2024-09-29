package graphs

func UndirectedPath(graph Graph[rune], nodeA, nodeB rune) bool {

	var hasPath func(rune, rune, map[rune]bool) bool
	hasPath = func(src, dst rune, visited map[rune]bool) bool {
		if src == dst {
			return true
		}
		if visited[src] {
			return false
		}
		visited[src] = true

		for _, key := range graph[src] {
			if hasPath(key, dst, visited) {
				return true
			}
		}

		return false
	}

	return hasPath(nodeA, nodeB, map[rune]bool{})
}
