package singlylinkedlist

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	expNode := &Node[string]{}
	expList := &List[string]{
		Head:   expNode,
		Tail:   expNode.Next,
		Length: 0,
	}

	list := New[string]()

	if !reflect.DeepEqual(expList, list) {
		t.Fatalf("Expected list: %v\nGot: %v", expList, list)
	}
}

func TestAppend(t *testing.T) {
	t.Parallel()

	expList := &List[int8]{}
	list := New[int8]()
	list.Head = nil

	list.Append(1)

	if !reflect.DeepEqual(expList, list) {
		t.Fatalf("Expected list: %v\nGot: %v", expList, list)
	}

	list = New[int8]()
	list.Append(15)
	list.Append(15)

	if list.Tail == nil || list.Length == 0 {
		t.Fatal("Expected list (length > 0, tail != nil)\nGot: ", list)
	}
}

func TestInsert(t *testing.T) {
	t.Parallel()

	// for coverage
	var node *Node[any]
	node.Insert(1)

	expList := New[int]()
	FillWithRange(expList, 1, 3)

	list := &Node[int]{
		Value: 1,
		Next: &Node[int]{
			Value: 3,
			Next:  nil,
		},
	}

	list.Insert(2)

	ptr1, ptr2 := list, expList.Head
	for ptr1 != nil || ptr2 != nil {
		if ptr1.Value != ptr2.Value {
			t.Fatalf("Expected value: %v\nGot: %v", ptr1.Value, ptr2.Value)
		}
		ptr1, ptr2 = ptr1.Next, ptr2.Next
	}
}

func TestAddNode(t *testing.T) {
	t.Parallel()

	expNode := &Node[string]{Value: "string"}
	node := AddNode("string")

	if !reflect.DeepEqual(expNode, node) {
		t.Fatalf("Expected node: %v\nGot: %v", expNode, node)
	}
}

func TestFillWithRange(t *testing.T) {
	t.Parallel()

	tail := &Node[int]{Value: 3}
	expList := &List[int]{
		Head: &Node[int]{
			Value: 1,
			Next: &Node[int]{
				Value: 2,
				Next:  tail,
			},
		},
		Tail:   tail,
		Length: 3,
	}

	list := New[int]()
	FillWithRange(list, 1, 3)

	if !reflect.DeepEqual(expList, list) {
		t.Fatalf("Expected list: %v\nGot: %v", expList, list)
	}
}

func TestFillWithInts(t *testing.T) {
	t.Parallel()

	list := New[int]()
	FillWithInts(list, []int{})

	if list.Length != 0 || list.Tail != nil {
		t.Fatal("Expected list(length > 0, tail != nil)\nGot:", list)
	}

	tail := &Node[int]{Value: 3}
	expList := &List[int]{
		Head: &Node[int]{
			Value: 1,
			Next: &Node[int]{
				Value: 2,
				Next:  tail,
			},
		},
		Tail:   tail,
		Length: 3,
	}

	list = New[int]()
	FillWithInts(list, []int{1, 2, 3})

	if !reflect.DeepEqual(expList, list) {
		t.Fatalf("Expected list: %v\nGot: %v", expList, list)
	}
}

func TestFillWithStrings(t *testing.T) {
	t.Parallel()

	list := New[string]()
	FillWithStrings(list)

	if list.Length != 0 || list.Tail != nil {
		t.Fatal("Expected list(length > 0, tail != nil)\nGot:", list)
	}

	tail := &Node[string]{Value: "3"}
	expList := &List[string]{
		Head: &Node[string]{
			Value: "1",
			Next: &Node[string]{
				Value: "2",
				Next:  tail,
			},
		},
		Tail:   tail,
		Length: 3,
	}

	list = New[string]()
	FillWithStrings(list, "1", "2", "3")

	if !reflect.DeepEqual(expList, list) {
		t.Fatalf("Expected list: %v\nGot: %v", expList, list)
	}
}
