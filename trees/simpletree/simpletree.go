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

func (t *Tree[T]) PrintAll() {

	fmt.Printf("\nlen:%d Depth:%d\nroot Node:\n %v", t.Len, t.Depth, t.Root.Data)

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

func (n *Node[T]) AddNode(data T) *Node[T] {
	n.Childrens = append(
		n.Childrens,
		&Node[T]{Data: data},
	)
	return n.Childrens[len(n.Childrens)-1]
}

/*
	// making Tree with Root Node and string "Root" in Data
	t := newTree("Root")

	// add Node to Root
	children1 := t.Root.addNode("children1")
	// inc Depth
	t.Depth++
	// inc Len
	t.Len++

	// add nodes to children1(Root-children1)
	children1.addNode("children1-1")
	t.Depth++
	t.Len++

	children1.addNode("children1-2")
	t.Len++

	children1.addNode("children1-3")
	t.Len++

	// add 2 more child to Root
	t.Root.addNode("children2")
	t.Len++
	children3 := t.Root.addNode("children3")
	t.Len++

	// add 2 child to children3(Root-children3)
	children3.addNode("children3-1")
	t.Len++
	children3.addNode("children3-2")
	t.Len++

	children4 := children3.addNode("children3-2")
	t.Len++

	children4.addNode("children3-2-1")
	t.Len++
	t.Depth++

	t.printAll()
*/
