package main

import "fmt"

// TODO: !!!!!!!
//package binarytree

// binaryTree represents a binary tree
// that holds values of any type.
type binaryTree[T any] struct {
	root       *node[T]
	len, depth int
}

// node represents a binary tree node
// that holds values of any type.
type node[T any] struct {
	data  T
	left  *node[T]
	right *node[T]
}

func newTree[T any](rootData T) binaryTree[T] {
	return binaryTree[T]{
		root: &node[T]{
			data: rootData,
		},
		len:   1,
		depth: 0,
	}
}

func generateBTree(a []int) binaryTree[int] {
	// get median index of array
	// set it to tree root
	m := medianIndex(len(a))
	t := newTree(a[m])
	f, s := a[:m], a[m+1:]

	if len(f) > len(s) {
		t.depth = calcDepth(f)
	}
	t.depth = calcDepth(s)

	t.root.left = generateBNode(f, &t.len)
	t.root.right = generateBNode(s, &t.len)
	return t
}

func generateBNode(a []int, length *int) *node[int] {
	m := medianIndex(len(a))
	if m == -1 {
		return nil
	}
	*length++
	return &node[int]{
		data:  a[m],
		left:  generateBNode(a[:m], length),
		right: generateBNode(a[m+1:], length),
	}
}

// maybe incorrect
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

func medianIndex(length int) int {
	return (length+1)/2 - 1
}

func generateSlice(from, to int) []int {
	var a []int
	for i := from; i <= to; i++ {
		a = append(a, i)
	}
	return a
}

// TODO tree balancing method
// нужен при добавлении к дереву элемента
// оптимизация узлов дерева для наиболее
// оптимальной расстановки

// TODO print tree func (not sure bout correct depth calc)
// на основе обхода дерева

func main() {
	// creating new binary tree with range from 1 to 10
	t := generateBTree(generateSlice(1, 12))

	fmt.Println(t.root.data)
	fmt.Println(t.len)
	fmt.Println(t.depth)

	//fmt.Println(t.root.left.data, "<--", t.root.data, "-->", t.root.right.data)
	//fmt.Println("--------------------")
	//
	//fmt.Println(t.root.left.left.data, "<--", t.root.left.data, "-->", t.root.left.right.data)
	//
	//fmt.Println("--------------------")
	//fmt.Println(t.root.right.left.data, "<--", t.root.right.data, "-->", t.root.right.right.data)
	//
	//fmt.Println("--------------------")
	//fmt.Println(t.root.right.right.right, "<--", t.root.right.right.data, "-->", t.root.right.right.left)

}
