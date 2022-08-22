package doublylinkedlist

import (
	"fmt"
)

// doublyLinkedList represents a linked list
// that holds values of any type.
type doublyLinkedList[T any] struct {
	head *node[T]
	len  int
	tail *node[T]
}

// node represents a doubly-linked node
// that holds values of any type.
type node[T any] struct {
	prev *node[T]
	val  T
	next *node[T]
}

func (l *doublyLinkedList[T]) append(v T) {
	if l.head == nil {
		return
	}
	ptr := l.head
	for {
		if ptr.next == nil {
			ptr.next = addNode(ptr, nil, v)
			break
		}
		ptr = ptr.next
	}
	l.tail = ptr.next
	l.len++
}

func (l *node[T]) insert(v T) {
	if l == nil {
		return
	}
	l.next = addNode[T](l, l.next, v)
	l.next.next.prev = l.next
}

func (l doublyLinkedList[T]) printList() {
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

func (l *node[T]) printNode() {
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

func (l doublyLinkedList[T]) printListReversed() {
	ptr := l.tail
	for {
		ptr.printNodeReversed()
		if ptr.prev == nil {
			break
		}
		ptr = ptr.prev
	}
	fmt.Println("len:", l.len)
}

func (l *node[T]) printNodeReversed() {
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

func newList[T any]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{
		head: &node[T]{},
		len:  0,
		tail: &node[T]{},
	}
}

func addNode[T any](prev, next *node[T], value T) *node[T] {
	return &node[T]{
		prev: prev,
		val:  value,
		next: next,
	}
}

func fillWithRange(l *doublyLinkedList[int], from, to int) {
	l.head.val = from
	l.len++
	ptr := l.head
	for i := from + 1; i <= to; i++ {
		if ptr.next == nil {
			ptr.next = addNode[int](ptr, nil, i)
			l.tail = ptr.next
			l.len++
		}
		ptr = ptr.next
	}
}

func fillWithStrings(l *doublyLinkedList[string], s ...string) {
	l.head.val = s[0]
	l.len++
	ptr := l.head
	for _, el := range s[1:] {
		if ptr.next == nil {
			ptr.next = addNode[string](ptr, nil, el)
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

	//add 22 between 3 and 4
	intList.head.next.next.insert(22)
	intList.len++
	intList.printList()

	// delete 22 between 3 and 4
	// need to know both 3 and 4 pointers to perform this
	// otherwise we can delete whole list after selected node
	//intList.head.next.next.next = nil
	tmp := intList.head.next.next
	intList.head.next.next.next = intList.head.next.next.next.next
	intList.head.next.next.next.prev = tmp
	intList.len--
	intList.printList()
	intList.printListReversed()

	fmt.Println("--------------")

	strList := newList[string]()
	fillWithStrings(strList, []string{"node1", "node2", "node3", "node4", "node5", "node6"}...)
	strList.printList()

	// append "node7" to the end of the list
	strList.append("node7")
	strList.printList()

	// add "inserted" between node3 and node4
	strList.head.next.next.insert("inserted")
	strList.len++
	strList.printList()

	// delete "inserted" between "node3" and "node4"
	// need to know both "node3" and "node4" pointers to perform this
	// otherwise we can delete whole list after selected node
	// strList.head.next.next = nil
	tmpStr := strList.head.next.next
	strList.head.next.next.next = strList.head.next.next.next.next
	strList.head.next.next.next.prev = tmpStr

	strList.len--
	strList.printList()
	strList.printListReversed()
*/
