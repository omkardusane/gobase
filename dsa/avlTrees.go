package dsa

type AVLTreeNode[T any] struct {
	value  T
	index  int
	height int
	left   *AVLTreeNode[T]
	right  *AVLTreeNode[T]
}

type AVLTree[T any] struct {
	root *AVLTreeNode[T]
}

func CreateAVLTree[T any]() *AVLTree[T] {
	return &AVLTree[T]{root: nil}
}

func (node *AVLTreeNode[T]) heightOrZero() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (tree *AVLTree[T]) Insert(value T, index int) {
	tree.root = tree.root.insert(value, index)
}

func (node *AVLTreeNode[T]) insert(value T, index int) *AVLTreeNode[T] {
	if node == nil {
		return &AVLTreeNode[T]{value: value, index: index, height: 1}
	}
	if index < node.index {
		node.left = node.left.insert(value, index)
	} else if index > node.index {
		node.right = node.right.insert(value, index)
	} else {
		// Index already exists, update value
		node.value = value
		return node
	}

	node.updateHeight()
	return node.rebalance()
}

func (node *AVLTreeNode[T]) updateHeight() {
	leftHeight := node.left.heightOrZero()
	rightHeight := node.right.heightOrZero()
	node.height = max(leftHeight, rightHeight) + 1
}

func (node *AVLTreeNode[T]) balanceFactor() int {
	return node.left.heightOrZero() - node.right.heightOrZero()
}

func (node *AVLTreeNode[T]) rebalance() *AVLTreeNode[T] {
	balanceFactor := node.balanceFactor()
	// Left-heavy
	if balanceFactor > 1 {
		return node.rotateRight()
	}
	// Right-heavy
	if balanceFactor < -1 {
		return node.rotateLeft()
	}
	return node
}

func (node *AVLTreeNode[T]) rotateLeft() *AVLTreeNode[T] {

}

func (node *AVLTreeNode[T]) rotateRight() *AVLTreeNode[T] {

}
