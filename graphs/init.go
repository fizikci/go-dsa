package graphs

import "fmt"

type Graph[T comparable] map[T][]T

var (
	SampleDAG   Graph[rune]
	SampleGraph Graph[rune]
)

func BuildGraph[T comparable](edges [][2]T, directed bool) Graph[T] {
	var graph Graph[T] = Graph[T]{}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		if !directed {
			graph[e[1]] = append(graph[e[1]], e[0])
		}
	}

	fmt.Println(graph)

	return graph
}

func (g Graph[T]) String() string {
	res := ""

	for key := range g {
		for _, e := range g[key] {
			res += fmt.Sprintf("%c => %c\n", key, e)
		}
	}

	return res
}

func init() {
	SampleDAG = BuildGraph[rune]([][2]rune{
		{'f', 'g'},
		{'f', 'i'},
		{'g', 'h'},
		{'i', 'g'},
		{'i', 'k'},
		{'j', 'i'},
	}, true)

	SampleGraph = BuildGraph([][2]rune{{'i', 'j'},
		{'k', 'i'},
		{'m', 'k'},
		{'k', 'l'},
		{'o', 'n'},
	}, false)
}
