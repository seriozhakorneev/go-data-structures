package singlylinkedlist

import (
	"fmt"
)

// SinglyLinkedList represents a linked list
// that holds values of any type.
type SinglyLinkedList[T any] struct {
	Head *Node[T]
	Len  int
}

// Node represents a singly-linked Node
// that holds values of any type.
type Node[T any] struct {
	Val  T
	Next *Node[T]
}

func (l *SinglyLinkedList[T]) Append(v T) {
	if l.Head == nil {
		return
	}
	ptr := l.Head
	for {
		if ptr.Next == nil {
			ptr.Next = addNode(v)
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
	l.Next = addNode[T](v)
	l.Next.Next = tmp
}

func (l *SinglyLinkedList[T]) PrintList() {
	ptr := l.Head
	for {
		ptr.printNode()
		if ptr.Next == nil {
			break
		}
		ptr = ptr.Next
	}
	fmt.Println("Len:", l.Len)
}

func (l *Node[T]) printNode() {
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

func NewList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Head: &Node[T]{},
		Len:  0,
	}
}

func addNode[T any](value T) *Node[T] {
	return &Node[T]{Val: value}
}

func FillWithRange(l *SinglyLinkedList[int], from, to int) {
	l.Head.Val = from
	l.Len++
	ptr := l.Head
	for i := from + 1; i <= to; i++ {
		if ptr.Next == nil {
			ptr.Next = addNode[int](i)
			l.Len++
		}
		ptr = ptr.Next
	}
}

func FillWithStrings(l *SinglyLinkedList[string], s ...string) {
	l.Head.Val = s[0]
	l.Len++
	ptr := l.Head
	for _, el := range s[1:] {
		if ptr.Next == nil {
			ptr.Next = addNode[string](el)
			l.Len++
		}
		ptr = ptr.Next
	}
}

/*
	intList := newList[int]()
	fillWithRange(intList, 1, 10)
	intList.printList()

	// append 11 to the end of the list
	intList.append(11)
	intList.printList()

	// add 22 between 3 and 4
	intList.Head.Next.Next.insert(22)
	intList.Len++
	intList.printList()

	// delete 22 between 3 and 4
	// need to know both 3 and 4 pointers to perform this
	// otherwise we can delete whole list after selected Node
	// intList.Head.Next.Next = nil
	intList.Head.Next.Next = intList.Head.Next.Next.Next
	intList.Len--
	intList.printList()

	fmt.Println("--------------")

	strList := newList[string]()
	fillWithStrings(strList, []string{"node1", "node2", "node3", "node4", "node5", "node6"}...)
	strList.printList()

	// append "node7" to the end of the list
	strList.append("node7")
	strList.printList()

	// add "inserted" between node3 and node4
	strList.Head.Next.insert("inserted")
	strList.Len++
	strList.printList()

	// delete "inserted" between "node3" and "node4"
	// need to know both "node3" and "node4" pointers to perform this
	// otherwise we can delete whole list after selected Node
	// strList.Head.Next.Next = nil
	strList.Head.Next.Next = strList.Head.Next.Next.Next
	strList.Len--
	strList.printList()
*/
