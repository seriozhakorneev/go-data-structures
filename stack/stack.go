package stack

import "fmt"

type Stack[T any] struct {
	St            []T
	Capacity, Len int
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf(
		"%v, Len(%v), cap(%v)",
		s.St, s.Len, s.Capacity,
	)
}

// New provide 0 Capacity to make Stack Capacity infinite
func New[T any](capacity int) Stack[T] {
	return Stack[T]{Capacity: capacity}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Len == 0
}

func (s *Stack[T]) IsFull() bool {
	if s.Len < s.Capacity || s.Capacity == 0 {
		return false
	}
	return true
}

func (s *Stack[T]) Size() int {
	return s.Len
}

func (s *Stack[T]) Top() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.St[len(s.St)-1], true
}

func (s *Stack[T]) Push(element T) bool {
	if !s.IsFull() {
		s.St = append(s.St, element)
		s.Len++
		return true
	}
	return false
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	element := s.St[len(s.St)-1]
	s.St = (s.St)[:len(s.St)-1]
	s.Len--
	return element, true
}
