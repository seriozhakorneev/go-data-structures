package main

// TODO: !!!!!!!
//package binarytree

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

func GenerateFromRange(from, to int) Tree[int] {
	a := generateSlice(from, to)
	// get median index of array
	// set it to tree Root
	m := medianIndex(len(a))
	t := NewTree(a[m])
	f, s := a[:m], a[m+1:]

	if len(f) > len(s) {
		t.Depth = calcDepth(f)
	}
	t.Depth = calcDepth(s)

	t.Root.Left = generateNode(f, &t.Len)
	t.Root.Right = generateNode(s, &t.Len)
	return t
}

func GenerateFromSlice(a []int) Tree[int] {
	// get median index of array
	// set it to tree Root
	m := medianIndex(len(a))
	t := NewTree(a[m])
	f, s := a[:m], a[m+1:]

	if len(f) > len(s) {
		t.Depth = calcDepth(f)
	}
	t.Depth = calcDepth(s)

	t.Root.Left = generateNode(f, &t.Len)
	t.Root.Right = generateNode(s, &t.Len)
	return t
}

func generateNode(a []int, length *int) *Node[int] {
	m := medianIndex(len(a))
	if m == -1 {
		return nil
	}
	*length++
	return &Node[int]{
		Data:  a[m],
		Left:  generateNode(a[:m], length),
		Right: generateNode(a[m+1:], length),
	}
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

//TODO maybe incorrect
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

// TODO tree balancing method
// нужен при добавлении к дереву элемента
// оптимизация узлов дерева для наиболее
// оптимальной расстановки

// TODO print tree func (not sure bout correct Depth calc)
// на основе обхода дерева

//func (t *Tree[T]) PrintAll() {
//
//	fmt.Printf("\nLength:%d Depth:%d Root Node:%v\n", t.Len, t.Depth, t.Root.Data)
//
//	var printRec func(level int, node *Node[T])
//	printRec = func(level int, node *Node[T]) {
//
//		if node == nil {
//			return
//		}
//		fmt.Print(" ", node.Data)
//		printRec(level+1, node.Right)
//		printRec(level+1, node.Left)
//	}
//
//	if t.Depth > 0 {
//		printRec(1, t.Root.Right)
//		printRec(1, t.Root.Left)
//	}
//}

func main() {
	// creating new binary tree with range from 1 to 10
	//t := NewTree("F")
	//
	//t.Root.Left = &Node[string]{Data: "B"}
	//t.Root.Left.Left = &Node[string]{Data: "A"}
	//t.Root.Left.Right = &Node[string]{Data: "D"}
	//t.Root.Left.Right.Left = &Node[string]{Data: "C"}
	//t.Root.Left.Right.Right = &Node[string]{Data: "E"}
	//
	//t.Root.Right = &Node[string]{Data: "G"}
	//t.Root.Right.Right = &Node[string]{Data: "I"}
	//t.Root.Right.Right.Left = &Node[string]{Data: "H"}
	//
	//fmt.Println(t.Root.Left.Data, "<--", t.Root.Data, "-->", t.Root.Right.Data)
	//fmt.Println("--------------------")
	//
	//fmt.Println(t.Root.Left.Left.Data, "<--", t.Root.Left.Data, "-->", t.Root.Left.Right.Data)
	//
	//fmt.Println("--------------------")
	//fmt.Println(t.Root.Right.Left, "<--", t.Root.Right.Data, "-->", t.Root.Right.Right.Data)
	//
	//fmt.Println("--------------------")
	//fmt.Println(t.Root.Right.Right.Left.Data, "<--", t.Root.Right.Right.Data, "-->", t.Root.Right.Right.Right)

	// LEFT
	// VISIT
	// RIGHT

}
