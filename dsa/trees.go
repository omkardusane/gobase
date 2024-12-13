package dsa

type GTreeNode[T any] struct {
	value T
	index int
	left  *GTreeNode[T]
	right *GTreeNode[T]
}

type BinaryTree[T any] struct {
	base  *GTreeNode[T]
	depth int
}

func CreateGTreeNode[T any]() *GTreeNode[T] {
	return &GTreeNode[T]{
		left:  nil,
		right: nil,
	}
}

func CreateBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{depth: 0, base: CreateGTreeNode[T]()}
}

func (tree *BinaryTree[T]) Traverse(reader func(nodeValue T, idx int)) {
	tree.base.Traverse(reader)
}

func (cursor *GTreeNode[T]) Traverse(reader func(nodeValue T, idx int)) {
	// in-order dfs
	if cursor.left != nil {
		cursor.left.Traverse(reader)
	}
	reader(cursor.value, cursor.index)
	if cursor.right != nil {
		cursor.right.Traverse(reader)
	}
}

func (tree *BinaryTree[T]) BfsTraverse(reader func(nodeValue T, depth int)) {
	currentDepth := 0
	maxDepth := 100
	reader(tree.base.value, currentDepth)
	tree.base.BfsTraverse(currentDepth+1, maxDepth, reader)
}

func (cursor *GTreeNode[T]) BfsTraverse(currentDepth int, maxDepth int, reader func(nodeValue T, depth int)) {
	thisDepth := CreateGList[*GTreeNode[T]]()
	if cursor.left != nil {
		reader(cursor.left.GetValue(), currentDepth)
		thisDepth.Push(cursor.left)
	}
	if cursor.right != nil {
		reader(cursor.right.GetValue(), currentDepth)
		thisDepth.Push(cursor.right)
	}
	currentDepth = currentDepth + 1
	// if maxDepth == currentDepth {
	// 	return
	// }
	thisDepth.ForEach(func(deeperCursor *GTreeNode[T], idx int) {
		deeperCursor.BfsTraverse(currentDepth, maxDepth, reader)
	})
}

func (tree *BinaryTree[T]) Find(index int) *GTreeNode[T] {
	// var cursor *GTreeNode[T] = tree.base;
	return tree.base.Find(index)
}

func (node *GTreeNode[T]) Find(index int) *GTreeNode[T] {
	if node.index == index {
		// matched
		return node
	} else {
		if node.left != nil {
			found := node.left.Find(index)
			if found != nil {
				return found
			}
		}
		if node.right != nil {
			found := node.right.Find(index)
			if found != nil {
				return found
			}
		}
		var nilVal *GTreeNode[T]
		return nilVal
	}
}

func (tree *BinaryTree[T]) Insert(value T, index int) {
	var cursor *GTreeNode[T] = tree.base
	cursor.Insert(value, index)
}

func (cursor *GTreeNode[T]) Insert(value T, index int) {
	var nilIdx int
	if cursor.index == nilIdx {
		cursor.value = value
		cursor.index = index
	} else {
		if cursor.index >= index {
			// we go left
			if cursor.left == nil {
				cursor.left = CreateGTreeNode[T]()
			}
			cursor.left.Insert(value, index)
		} else {
			// we go right
			if cursor.right == nil {
				cursor.right = CreateGTreeNode[T]()
			}
			cursor.right.Insert(value, index)
		}
	}
}

func (node *GTreeNode[T]) GetValue() T {
	return node.value
}
