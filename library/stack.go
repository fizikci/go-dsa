package library

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() T {
	removed := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return removed
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
