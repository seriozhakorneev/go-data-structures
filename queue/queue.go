package queue

import "fmt"

type Queue[T any] struct {
	Qu       []T
	Cap, Len int
}

func (q *Queue[T]) String() string {
	for i := q.Len - 1; i > -1; i-- {
		fmt.Print(q.Qu[i], " ")
	}

	return fmt.Sprintf(
		"Length(%v), cap(%v)",
		q.Len, q.Cap,
	)
}

// New - returns new Queue with provided capacity.
// provide 0 capacity to make Queue capacity infinite
func New[T any](capacity int) Queue[T] {
	return Queue[T]{Cap: capacity}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Len == 0
}

func (q *Queue[T]) IsFull() bool {
	if q.Len < q.Cap || q.Cap == 0 {
		return false
	}

	return true
}

func (q *Queue[T]) Size() int {
	return q.Len
}

func (q *Queue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	return q.Qu[0], true
}

func (q *Queue[T]) Enqueue(element T) bool {
	if !q.IsFull() {
		q.Qu = append(q.Qu, element)
		q.Len++
		return true
	}

	return false
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	element := q.Qu[0]
	q.Qu = (q.Qu)[1:]
	q.Len--

	return element, true
}
