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
	// emit this cursor
	reader(cursor.value, cursor.index)
	if cursor.left != nil {
		cursor.left.Traverse(reader)
	}
	if cursor.right != nil {
		cursor.right.Traverse(reader)
	}
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
