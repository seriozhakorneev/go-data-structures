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

func GenerateFromRange(from, to int) Tree[int] {
	a := generateSlice(from, to)
	// get median index of array
	// set it to tree Root
	m := medianIndex(len(a))
	t := NewTree(a[m])
	f, s := a[:m], a[m+1:]

	if len(f) > len(s) {
		t.Depth = CalcDepth(f)
	}
	t.Depth = CalcDepth(s)

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
		t.Depth = CalcDepth(f)
	}
	t.Depth = CalcDepth(s)

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

// CalcDepth maybe incorrect
func CalcDepth(a []int) int {
	if len(a) > 2 {
		depth := float64(len(a)) / float64(2)
		if depth/10 != 0 {
			depth++
		}
		return int(depth)
	}
	return len(a)
}

func (t *Tree[T]) FindDepth() int {
	var find func(node *Node[T]) int
	find = func(node *Node[T]) int {
		if node == nil {
			return 0
		}

		lDepth := find(node.Left)
		rDepth := find(node.Right)

		if lDepth > rDepth {
			return lDepth + 1
		}
		return rDepth + 1
	}

	t.Depth = find(t.Root)
	return t.Depth
}

// TODO tree balancing method
// нужен при добавлении к дереву элемента
// оптимизация узлов дерева для наиболее
// оптимальной расстановки
