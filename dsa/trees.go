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

func CreateBinaryTree[T any]() *BinaryTree[T] {
	return &BinaryTree[T]{depth: 0}
}

func (tree *BinaryTree[T]) Traverse(reader func(nodeValue T, idx int)) {
	tree.base.Traverse(reader)
}

func (cursor *GTreeNode[T]) Traverse(reader func(nodeValue T, idx int)) {
	// emit this cursor
	reader(cursor.value, cursor.index)
	if cursor.left != nil {
		cursor.left.Traverse(reader)
	} else if cursor.right != nil {
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

func (tree *BinaryTree[T]) Insert(index int, value T) {
	var cursor *GTreeNode[T]
	if cursor.index == index {
		// matched
	} else if cursor.left != nil {

	} else if cursor.right != nil {

	}
}
