package btree

import (
	"sort"
)

// Tree represents a b-tree
// that holds values of any type.
type Tree struct {
	Root       *Node
	Len, Depth int
}

// Node represents a b-tree Node
// that holds values of any type.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewTree(rootData int) Tree {
	return Tree{
		Root: &Node{
			Value: rootData,
		},
		Len:   1,
		Depth: 0,
	}
}

// AddNodes Add new Nodes with provided value/values
// triggers tree balancing method
func (t *Tree) AddNodes(values ...int) {
	t.Balancing(values...)
}

func GenFromRange(from, to int) Tree {
	a := genSlice(from, to)
	// get median index of array
	// set it to tree Root
	m := medianIndex(len(a))
	t := NewTree(a[m])
	f, s := a[:m], a[m+1:]

	if len(f) > len(s) {
		t.Depth = calcDepth(f)
	}
	t.Depth = calcDepth(s)

	t.Root.Left = genNode(f, &t.Len)
	t.Root.Right = genNode(s, &t.Len)
	return t
}

func GenFromSlice(a []int) Tree {
	if !sort.IntsAreSorted(a) {
		sort.Ints(a)
	}

	// get median index of array
	// set it to tree Root
	m := medianIndex(len(a))
	t := NewTree(a[m])
	f, s := a[:m], a[m+1:]

	if len(f) > len(s) {
		t.Depth = calcDepth(f)
	}
	t.Depth = calcDepth(s)

	t.Root.Left = genNode(f, &t.Len)
	t.Root.Right = genNode(s, &t.Len)
	return t
}

func genNode(a []int, length *int) *Node {
	m := medianIndex(len(a))
	if m == -1 {
		return nil
	}
	*length++
	return &Node{
		Value: a[m],
		Left:  genNode(a[:m], length),
		Right: genNode(a[m+1:], length),
	}
}

func medianIndex(length int) int {
	return (length+1)/2 - 1
}

func genSlice(from, to int) []int {
	var a []int
	for i := from; i <= to; i++ {
		a = append(a, i)
	}
	return a
}

func calcDepth(a []int) int {
	if len(a) > 2 {
		depth := float64(len(a)) / float64(2)
		if depth/10 != 0 {
			depth++
		}
		return int(depth)
	}
	return len(a)
}

// Balancing makes binary tree balanced b-tree
// values can be added to slice on which origin
// we will make new tree
func (t *Tree) Balancing(values ...int) {
	s := values

	s = travWithSort(t.Root, s...)

	m := medianIndex(len(s))
	newT := NewTree(s[m])
	f, s := s[:m], s[m+1:]

	if len(f) > len(s) {
		newT.Depth = calcDepth(f)
	}
	newT.Depth = calcDepth(s)

	newT.Root.Left = genNode(f, &newT.Len)
	newT.Root.Right = genNode(s, &newT.Len)

	*t = newT
}

// travWithSort InOrder traversal with easy sorting
func travWithSort(n *Node, values ...int) (a []int) {
	a = values

	var inOrder func(node *Node)
	inOrder = func(node *Node) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		a = append(a, node.Value)
		sortLast(a)
		inOrder(node.Right)
	}

	inOrder(n)
	return a
}

// GetSlice by returns slice a,
// collected by inOrder traversal
func (t *Tree) GetSlice() (a []int) {
	var inOrder func(node *Node)
	inOrder = func(node *Node) {
		if node == nil {
			return
		}

		inOrder(node.Left)
		a = append(a, node.Value)
		inOrder(node.Right)
	}

	inOrder(t.Root)
	return a
}

// sortLast sorting all elements from last to first
// sort s in increasing order
func sortLast(s []int) {
	for i := len(s) - 1; i > 0 && s[i-1] > s[i]; i-- {
		tmp := s[i]
		s[i] = s[i-1]
		s[i-1] = tmp
	}
}
