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
	Data  int
	Left  *Node
	Right *Node
}

func NewTree(rootData int) Tree {
	return Tree{
		Root: &Node{
			Data: rootData,
		},
		Len:   1,
		Depth: 0,
	}
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
		Data:  a[m],
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
func (t *Tree) Balancing() {
	// In order traversal
	a := inOrder(t.Root)
	// generate new tree
	*t = GenFromSlice(a)
}

func inOrder(root *Node) (stack []int) {
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}

		traverse(node.Left)
		stack = append(stack, node.Data)
		traverse(node.Right)
	}

	traverse(root)
	return stack
}
