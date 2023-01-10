package binarytree

// Tree represents a binary tree
// that holds values of any type.
type Tree[T any] struct {
	Root       *Node[T]
	Len, Depth int
}

// Node represents a binary tree Node
// that holds values of any type.
type Node[T any] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

func NewTree[T any](rootData T) Tree[T] {
	return Tree[T]{
		Root: &Node[T]{
			Data: rootData,
		},
		Len:   1,
		Depth: 0,
	}
}

func (n *Node[T]) AddLeft(val T) *Node[T] {
	left := &Node[T]{Data: val}
	n.Left = left
	return left
}

func (n *Node[T]) AddRight(val T) *Node[T] {
	right := &Node[T]{Data: val}
	n.Right = right
	return right
}

func (t *Tree[T]) CalcDepNLen() {
	var rec func(node *Node[T], depth int) int

	rec = func(node *Node[T], depth int) int {
		if node == nil {
			return depth
		}
		t.Len++
		depth++

		lDepth := rec(node.Left, depth)
		rDepth := rec(node.Right, depth)

		if lDepth > rDepth {
			return lDepth
		}
		return rDepth
	}

	t.Depth += rec(t.Root, t.Depth)
	t.Depth--
	t.Len--
}
