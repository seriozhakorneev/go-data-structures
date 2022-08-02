// TODO: !!!!!!!
package main

//package binarytree

// binaryTree represents a binary binaryTree
// that holds values of any type.
type binaryTree[T any] struct {
	root       *node[T]
	len, depth int
}

// node represents a binary binaryTree node
// that holds values of any type.
type node[T any] struct {
	left  *node[T]
	data  T
	right *node[T]
}

func newTree[T any](rootData T) binaryTree[T] {
	return binaryTree[T]{
		root: &node[T]{
			data: rootData,
		},
		len:   1,
		depth: 1,
	}
}

// TODO tree balancing method
// нужен при добавлении к дереву элемента
// оптимизация узлов дерева для наиболее
// оптимальной расстановки

func createWithRange(from, to int) binaryTree[int] {

	// generate
	//for i := from + 1; i <= to; i++ {
	//	fmt.Print(i, " ")
	//}

	// найти медианну и посетить в корень
	t := newTree(10)

	return t
}

func main() {
	// creating new binary tree with range from 1 to 50
	createWithRange(1, 50)
}
