package stack

import "fmt"

type stack[T any] struct {
	st            []T
	capacity, len int
}

func (s stack[T]) String() string {
	return fmt.Sprintf(
		"%v, len(%v), cap(%v)",
		s.st, s.len, s.capacity,
	)
}

// provide 0 capacity to make stack capacity infinite
func newStack[T any](capacity int) stack[T] {
	return stack[T]{capacity: capacity}
}

func (s stack[T]) isEmpty() bool {
	return s.len == 0
}

func (s *stack[T]) isFull() bool {
	if s.len < s.capacity || s.capacity == 0 {
		return false
	}
	return true
}

func (s stack[T]) size() int {
	return s.len
}

func (s stack[T]) top() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}

	return s.st[len(s.st)-1], true
}

func (s *stack[T]) push(element T) bool {
	if !s.isFull() {
		s.st = append(s.st, element)
		s.len++
		return true
	}
	return false
}

func (s *stack[T]) pop() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}

	element := s.st[len(s.st)-1]
	s.st = (s.st)[:len(s.st)-1]
	s.len--
	return element, true
}

// new stack with capacity 3
//st := newStack[string](3)
//fmt.Println("is empty:", st.isEmpty())
//fmt.Println(st)
//
//fmt.Println(
//	"pushing",
//	st.push("1"),
//	st.push("2"),
//	st.push("3"),
//	st.push("4"),
//)
//
//fmt.Println("is empty:", st.isEmpty())
//fmt.Println(st)
//
//top, _ := st.top()
//fmt.Println("top:", top)
//fmt.Println("size:", st.size())
//
//if el, ok := st.pop(); ok {
//	fmt.Println("pop el:", el)
//}
//if el, ok := st.pop(); ok {
//	fmt.Println("pop el:", el)
//}
//
//if el, ok := st.top(); ok {
//	fmt.Println("top:", el)
//}
//
//if el, ok := st.pop(); ok {
//	fmt.Println("pop el:", el)
//}
//
//fmt.Println("size:", st.size())
//fmt.Println(st.isEmpty(), st)
