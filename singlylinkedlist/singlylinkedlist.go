package singlylinkedlist

import (
	"fmt"
)

// SinglyLinkedList represents a linked list
// that holds values of any type.
type SinglyLinkedList[T any] struct {
	head *Node[T]
	len  int
}

// Node represents a singly-linked Node
// that holds values of any type.
type Node[T any] struct {
	val  T
	next *Node[T]
}

func (l *SinglyLinkedList[T]) Append(v T) {
	if l.head == nil {
		return
	}
	ptr := l.head
	for {
		if ptr.next == nil {
			ptr.next = addNode(v)
			break
		}
		ptr = ptr.next
	}
	l.len++
}

func (l *Node[T]) Insert(v T) {
	if l == nil {
		return
	}
	tmp := l.next
	l.next = addNode[T](v)
	l.next.next = tmp
}

func (l *SinglyLinkedList[T]) PrintList() {
	ptr := l.head
	for {
		ptr.printNode()
		if ptr.next == nil {
			break
		}
		ptr = ptr.next
	}
	fmt.Println("len:", l.len)
}

func (l *Node[T]) printNode() {
	if l == nil {
		fmt.Print(nil)
		return
	}
	fmt.Printf("%v->%v ",
		l.val,
		func() interface{} {
			if l.next != nil {
				return l.next.val
			}
			return nil
		}())
}

func NewList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: &Node[T]{},
		len:  0,
	}
}

func addNode[T any](value T) *Node[T] {
	return &Node[T]{val: value}
}

func FillWithRange(l *SinglyLinkedList[int], from, to int) {
	l.head.val = from
	l.len++
	ptr := l.head
	for i := from + 1; i <= to; i++ {
		if ptr.next == nil {
			ptr.next = addNode[int](i)
			l.len++
		}
		ptr = ptr.next
	}
}

func FillWithStrings(l *SinglyLinkedList[string], s ...string) {
	l.head.val = s[0]
	l.len++
	ptr := l.head
	for _, el := range s[1:] {
		if ptr.next == nil {
			ptr.next = addNode[string](el)
			l.len++
		}
		ptr = ptr.next
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
	intList.head.next.next.insert(22)
	intList.len++
	intList.printList()

	// delete 22 between 3 and 4
	// need to know both 3 and 4 pointers to perform this
	// otherwise we can delete whole list after selected Node
	// intList.head.next.next = nil
	intList.head.next.next = intList.head.next.next.next
	intList.len--
	intList.printList()

	fmt.Println("--------------")

	strList := newList[string]()
	fillWithStrings(strList, []string{"node1", "node2", "node3", "node4", "node5", "node6"}...)
	strList.printList()

	// append "node7" to the end of the list
	strList.append("node7")
	strList.printList()

	// add "inserted" between node3 and node4
	strList.head.next.insert("inserted")
	strList.len++
	strList.printList()

	// delete "inserted" between "node3" and "node4"
	// need to know both "node3" and "node4" pointers to perform this
	// otherwise we can delete whole list after selected Node
	// strList.head.next.next = nil
	strList.head.next.next = strList.head.next.next.next
	strList.len--
	strList.printList()
*/
