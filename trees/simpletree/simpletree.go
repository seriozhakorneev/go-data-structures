package simpletree

import "fmt"

// Tree represents a Tree
// that holds values of any type.
type Tree[T any] struct {
	Root       *Node[T]
	Len, Depth int
}

// Node represents a Tree Node
// that holds values of any type.
type Node[T any] struct {
	Data      T
	Childrens []*Node[T]
}

func NewTree[T any](rootData T) Tree[T] {
	return Tree[T]{
		Root:  &Node[T]{Data: rootData},
		Len:   1,
		Depth: 0,
	}
}

func (n *Node[T]) AddNode(data T) *Node[T] {
	n.Childrens = append(
		n.Childrens,
		&Node[T]{Data: data},
	)
	return n.Childrens[len(n.Childrens)-1]
}

func (t *Tree[T]) PrintAll() {
	fmt.Printf(
		"\nlen:%d Depth:%d\nroot Node:\n %v",
		t.Len,
		t.Depth,
		t.Root.Data,
	)

	var printRec func(level int, nodes []*Node[T])
	printRec = func(level int, nodes []*Node[T]) {

		fmt.Printf("\n%d level nodes:\n", level)
		var next []*Node[T]
		for _, n := range nodes {
			fmt.Print(" ", n.Data)
			next = append(next, n.Childrens...)
		}

		if level == t.Depth {
			return
		}
		printRec(level+1, next)
	}

	if t.Depth > 0 {
		printRec(1, t.Root.Childrens)
	}
}

func (t *Tree[T]) FindDepth() int {
	var find func(node *Node[T]) int
	find = func(node *Node[T]) int {
		if node == nil {
			return 0
		}

		maxDepth := 0
		for _, children := range node.Childrens {
			depth := find(children)
			if depth > maxDepth {
				maxDepth = depth
			}
		}

		return maxDepth + 1
	}

	t.Depth = find(t.Root)
	return t.Depth
}
