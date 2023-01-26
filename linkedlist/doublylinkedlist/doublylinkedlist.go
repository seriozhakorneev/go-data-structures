package doublylinkedlist

import "fmt"

// List represents a doubly-linked List
// that holds values of any type.
type List[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

// Node represents a doubly-linked Node
// that holds values of any type.
type Node[T any] struct {
	Prev  *Node[T]
	Value T
	Next  *Node[T]
}

// New - returns new List.
func New[T any]() *List[T] {
	head := &Node[T]{}

	return &List[T]{
		Head:   head,
		Length: 0,
		Tail:   head.Next,
	}
}

// Append - adds new Tail to List, after last Tail.
func (l *List[T]) Append(v T) {
	if l.Head == nil {
		return
	}

	ptr := l.Head
	tail := AddNode(ptr, nil, v)

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

// Insert - adds new Node, on position after current Node.
// If any Node exists next to current, it becomes next to new.
// Method disconnected with List type, will not affect Tail or Length.
func (l *Node[T]) Insert(v T) {
	if l == nil {
		return
	}

	l.Next = AddNode[T](l, l.Next, v)
	l.Next.Next.Prev = l.Next
}

// AddNode - returns new Node with provided parameters.
func AddNode[T any](prev, next *Node[T], value T) *Node[T] {
	return &Node[T]{
		Prev:  prev,
		Value: value,
		Next:  next,
	}
}

// FillWithRange - generate List from provided range.
func FillWithRange(l *List[int], from, to int) {
	l.Head.Value = from
	l.Length++

	ptr := l.Head
	for i := from + 1; i <= to; i++ {
		if ptr.Next == nil {
			ptr.Next = AddNode[int](ptr, nil, i)
			l.Tail = ptr.Next
			l.Length++
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
			node := AddNode[string](ptr, nil, el)
			ptr.Next = node
			l.Length++

			if i == len(s[1:])-1 {
				l.Tail = node
			}
		}
		ptr = ptr.Next
	}
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

// PrintNode - print Node Value, and its Next Node Value.
func (l *Node[T]) PrintNode() {
	if l == nil {
		fmt.Print(nil)
		return
	}

	fmt.Printf("%v<-%v->%v ",
		func() interface{} {
			if l.Prev != nil {
				return l.Prev.Value
			}
			return nil
		}(),
		l.Value,
		func() interface{} {
			if l.Next != nil {
				return l.Next.Value
			}
			return nil
		}(),
	)
}

// PrintListReversed - prints all Node's in reversed order, from Tail to Head.
func (l *List[T]) PrintListReversed() {
	ptr := l.Tail

	for {
		ptr.PrintNodeReversed()
		if ptr.Prev == nil {
			break
		}
		ptr = ptr.Prev
	}

	fmt.Println("Length:", l.Length)
}

// PrintNodeReversed - print Node Value, and its Next Node Value, in reversed order.
func (l *Node[T]) PrintNodeReversed() {
	if l == nil {
		fmt.Print(nil)
		return
	}

	fmt.Printf("%v<-%v->%v ",
		func() interface{} {
			if l.Next != nil {
				return l.Next.Value
			}
			return nil
		}(),
		l.Value,
		func() interface{} {
			if l.Prev != nil {
				return l.Prev.Value
			}
			return nil
		}(),
	)
}
