package singlylinkedlist

import (
	"fmt"
)

// List - represents a singly-linked list,
// that holds values of any type.
type List[T any] struct {
	// Head - first element of linked list.
	Head *Node[T]
	// Tail - last element of linked list.
	Tail   *Node[T]
	Length int
}

// Node - represents a singly-linked Node,
// that holds values of any type.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// NewList - returns new List.
func NewList[T any]() *List[T] {
	head := &Node[T]{}

	return &List[T]{
		Head:   head,
		Tail:   head.Next,
		Length: 0,
	}
}

// Append - adds new Tail to List, after last Tail.
func (l *List[T]) Append(v T) {
	if l.Head == nil {
		return
	}

	ptr, tail := l.Head, AddNode(v)
	for {
		if ptr.Next == nil {
			ptr.Next = tail
			break
		}
		ptr = ptr.Next
	}

	l.Tail = tail
	l.Length++
}

// InsertNext - adds new Node, on position after current Node.
// If any Node exists next to current, it becomes next to new.
func (l *Node[T]) InsertNext(v T) {
	if l == nil {
		return
	}

	tmp := l.Next
	l.Next = AddNode[T](v)
	l.Next.Next = tmp
}

// PrintList - prints all Node's, from Head to Tail.
func (l *List[T]) PrintList() {
	ptr := l.Head

	for {
		ptr.PrintNode()
		if ptr.Next == nil {
			break
		}
		ptr = ptr.Next
	}

	fmt.Println("Length:", l.Length)
}

// PrintNode - print Node Value, and its Next Node Value
func (l *Node[T]) PrintNode() {
	if l == nil {
		fmt.Print(nil)
		return
	}

	fmt.Printf("%v->%v ",
		l.Value,
		func() any {
			if l.Next != nil {
				return l.Next.Value
			}
			return nil
		}())
}

// AddNode - returns new Node with provided value.
func AddNode[T any](v T) *Node[T] {
	return &Node[T]{Value: v}
}

// FillWithRange - generate List from provided range.
func FillWithRange(l *List[int], from, to int) {
	l.Head.Value = from
	l.Length++

	ptr := l.Head
	for i := from + 1; i <= to; i++ {
		if ptr.Next == nil {
			node := AddNode[int](i)
			ptr.Next = node
			l.Length++

			if i == to {
				l.Tail = node
			}
		}
		ptr = ptr.Next
	}
}

// FillWithInts - generate List from provided slice.
func FillWithInts(l *List[int], a []int) {
	if len(a) == 0 {
		return
	}

	l.Head.Value = a[0]
	l.Length++

	ptr := l.Head
	for i, el := range a[1:] {
		if ptr.Next == nil {
			node := AddNode[int](el)
			ptr.Next = node
			l.Length++

			if i == len(a)-1 {
				l.Tail = node
			}
		}
		ptr = ptr.Next
	}
}

// FillWithStrings - generate List from provided strings.
func FillWithStrings(l *List[string], s ...string) {
	if len(s) == 0 {
		return
	}

	l.Head.Value = s[0]
	l.Length++

	ptr := l.Head
	for i, el := range s[1:] {
		if ptr.Next == nil {
			node := AddNode[string](el)
			ptr.Next = node
			l.Length++

			if i == len(s)-1 {
				l.Tail = node
			}
		}
		ptr = ptr.Next
	}
}
