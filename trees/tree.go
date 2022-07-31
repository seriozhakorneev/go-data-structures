//package trees
package main

import "fmt"

// tree represents a tree
// that holds values of any type.
type tree[T any] struct {
	root       *node[T]
	len, depth int
}

// node represents a tree node
// that holds values of any type.
type node[T any] struct {
	data      T
	childrens []*node[T]
}

func (n *node[T]) addNode(data T) *node[T] {
	n.childrens = append(
		n.childrens,
		&node[T]{data: data},
	)
	return n.childrens[len(n.childrens)-1]
}

func newTree[T any](rootData T) tree[T] {
	return tree[T]{
		root:  &node[T]{data: rootData},
		len:   1,
		depth: 1,
	}
}

func (t *tree[T]) printAll() {

	fmt.Printf("\nlen:%d depth:%d\n%d level nodes:\n %v", t.len, t.depth, 1, t.root.data)

	var printRecursively func(level int, nodes []*node[T])
	printRecursively = func(level int, nodes []*node[T]) {

		fmt.Printf("\n%d level nodes:\n", level)
		var next []*node[T]
		for _, n := range nodes {
			fmt.Print(" ", n.data)
			next = append(next, n.childrens...)
		}

		if level == t.depth {
			return
		}
		printRecursively(level+1, next)
	}

	if t.depth > 1 {
		printRecursively(2, t.root.childrens)
	}
}

/*
	// making tree with root node and string "root" in data
	t := newTree("root")

	// add node to root
	children1 := t.root.addNode("children1")
	// inc depth
	t.depth++
	// inc len
	t.len++

	// add nodes to children1(root-children1)
	children1.addNode("children1-1")
	t.depth++
	t.len++

	children1.addNode("children1-2")
	t.len++

	children1.addNode("children1-3")
	t.len++

	// add 2 more child to root
	t.root.addNode("children2")
	t.len++
	children3 := t.root.addNode("children3")
	t.len++

	// add 2 child to children3(root-children3)
	children3.addNode("children3-1")
	t.len++
	children3.addNode("children3-2")
	t.len++

	children4 := children3.addNode("children3-2")
	t.len++

	children4.addNode("children3-2-1")
	t.len++
	t.depth++

	t.printAll()
*/
