package singlylinkedlist

import (
	"fmt"
)

// List represents a singly linked list
// that holds values of any type.
type List[T any] struct {
	Head *Node[T]
	Len  int
}

// Node represents a singly-linked Node
// that holds values of any type.
type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (l *List[T]) Append(v T) {
	if l.Head == nil {
		return
	}
	ptr := l.Head
	for {
		if ptr.Next == nil {
			ptr.Next = AddNode(v)
			break
		}
		ptr = ptr.Next
	}
	l.Len++
}

func (l *Node[T]) Insert(v T) {
	if l == nil {
		return
	}
	tmp := l.Next
	l.Next = AddNode[T](v)
	l.Next.Next = tmp
}

func (l *List[T]) PrintList() {
	ptr := l.Head
	for {
		ptr.PrintNode()
		if ptr.Next == nil {
			break
		}
		ptr = ptr.Next
	}
	fmt.Println("Len:", l.Len)
}

func (l *Node[T]) PrintNode() {
	if l == nil {
		fmt.Print(nil)
		return
	}
	fmt.Printf("%v->%v ",
		l.Val,
		func() interface{} {
			if l.Next != nil {
				return l.Next.Val
			}
			return nil
		}())
}

func NewList[T any]() *List[T] {
	return &List[T]{
		Head: &Node[T]{},
		Len:  0,
	}
}

func AddNode[T any](value T) *Node[T] {
	return &Node[T]{Val: value}
}

func FillWithRange(l *List[int], from, to int) {
	l.Head.Val = from
	l.Len++
	ptr := l.Head
	for i := from + 1; i <= to; i++ {
		if ptr.Next == nil {
			ptr.Next = AddNode[int](i)
			l.Len++
		}
		ptr = ptr.Next
	}
}

func FillWithInts(l *List[int], a []int) {
	if len(a) == 0 {
		return
	}
	l.Head.Val = a[0]
	l.Len++
	ptr := l.Head
	for _, el := range a[1:] {
		if ptr.Next == nil {
			ptr.Next = AddNode[int](el)
			l.Len++
		}
		ptr = ptr.Next
	}
}

func FillWithStrings(l *List[string], s ...string) {
	if len(s) == 0 {
		return
	}
	l.Head.Val = s[0]
	l.Len++
	ptr := l.Head
	for _, el := range s[1:] {
		if ptr.Next == nil {
			ptr.Next = AddNode[string](el)
			l.Len++
		}
		ptr = ptr.Next
	}
}
