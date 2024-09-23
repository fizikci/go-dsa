package library

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}
func (q *Queue[T]) Dequeue() T {
	res := q.items[0]
	q.items = q.items[1:]
	return res
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
