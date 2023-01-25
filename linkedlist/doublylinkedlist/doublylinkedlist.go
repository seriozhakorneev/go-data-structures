package doublylinkedlist

import (
	"fmt"
)

// List represents a doubly linked List
// that holds values of any type.
type List[T any] struct {
	head *Node[T]
	len  int
	tail *Node[T]
}

// Node represents a doubly-linked Node
// that holds values of any type.
type Node[T any] struct {
	prev *Node[T]
	val  T
	next *Node[T]
}

func (l *List[T]) Append(v T) {
	if l.head == nil {
		return
	}
	ptr := l.head
	for {
		if ptr.next == nil {
			ptr.next = AddNode(ptr, nil, v)
			break
		}
		ptr = ptr.next
	}
	l.tail = ptr.next
	l.len++
}

func (l *Node[T]) Insert(v T) {
	if l == nil {
		return
	}
	l.next = AddNode[T](l, l.next, v)
	l.next.next.prev = l.next
}

func (l *List[T]) PrintList() {
	ptr := l.head
	for {
		ptr.Print()
		if ptr.next == nil {
			break
		}
		ptr = ptr.next
	}
	fmt.Println("len:", l.len)
}

func (l *Node[T]) Print() {
	if l == nil {
		fmt.Print(nil)
		return
	}
	fmt.Printf("%v<-%v->%v ",
		func() interface{} {
			if l.prev != nil {
				return l.prev.val
			}
			return nil
		}(),
		l.val,
		func() interface{} {
			if l.next != nil {
				return l.next.val
			}
			return nil
		}(),
	)
}

func (l *List[T]) PrintListReversed() {
	ptr := l.tail
	for {
		ptr.PrintNodeReversed()
		if ptr.prev == nil {
			break
		}
		ptr = ptr.prev
	}
	fmt.Println("len:", l.len)
}

func (l *Node[T]) PrintNodeReversed() {
	if l == nil {
		fmt.Print(nil)
		return
	}
	fmt.Printf("%v<-%v->%v ",
		func() interface{} {
			if l.next != nil {
				return l.next.val
			}
			return nil
		}(),
		l.val,
		func() interface{} {
			if l.prev != nil {
				return l.prev.val
			}
			return nil
		}(),
	)
}

func New[T any]() *List[T] {
	return &List[T]{
		head: &Node[T]{},
		len:  0,
		tail: &Node[T]{},
	}
}

func AddNode[T any](prev, next *Node[T], value T) *Node[T] {
	return &Node[T]{
		prev: prev,
		val:  value,
		next: next,
	}
}

func FillWithRange(l *List[int], from, to int) {
	l.head.val = from
	l.len++
	ptr := l.head
	for i := from + 1; i <= to; i++ {
		if ptr.next == nil {
			ptr.next = AddNode[int](ptr, nil, i)
			l.tail = ptr.next
			l.len++
		}
		ptr = ptr.next
	}
}

func FillWithStrings(l *List[string], s ...string) {
	l.head.val = s[0]
	l.len++
	ptr := l.head
	for _, el := range s[1:] {
		if ptr.next == nil {
			ptr.next = AddNode[string](ptr, nil, el)
			l.len++
		}
		ptr = ptr.next
	}
}
