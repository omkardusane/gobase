package dsa

import "fmt"

type GTreeNode[T any] struct {
	value  T
	index  int
	depth  int
	Height int
	left   *GTreeNode[T]
	right  *GTreeNode[T]
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
	// inorder traversal
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

func (tree *BinaryTree[T]) BfsTraverse(reader func(nodeValue T, currentDepth int)) {
	currentDepth := 0
	maxDepth := 100
	reader(tree.base.value, currentDepth)
	tree.base.BfsTraverse(currentDepth+1, maxDepth, reader)
}

func (cursor *GTreeNode[T]) BfsTraverse(currentDepth int, maxDepth int, reader func(nodeValue T, depth int)) {
	thisDepth := CreateGList[*GTreeNode[T]]()
	if cursor.left != nil {
		fmt.Print("node ", cursor.GetValue())
		fmt.Print("-child-Left-", cursor.left.GetValue())
		reader(cursor.left.GetValue(), currentDepth)
		thisDepth.Push(cursor.left)
	}
	if cursor.right != nil {
		fmt.Print("node ", cursor.GetValue())
		fmt.Print("-child-Right-", cursor.right.GetValue())
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

func (tree *BinaryTree[T]) GetRoot() *GTreeNode[T] {
	return tree.base
}

func (tree *BinaryTree[T]) FindElemAt(idx int) T {
	cursor := tree.base
	return cursor.FindElemAt(idx)
}

func (cursor *GTreeNode[T]) FindElemAt(idx int) T {
	var nilVal int
	var nilT T
	if cursor.index == nilVal {
		return nilT
	}
	if cursor.index == idx {
		return cursor.value
	} else if cursor.index > idx && cursor.left != nil {
		//  go find in left subtree
		return cursor.left.FindElemAt(idx)
	} else if cursor.index < idx && cursor.right != nil {
		//  go find in right subtree
		return cursor.right.FindElemAt(idx)
	}
	return nilT
}

func (cursor *GTreeNode[T]) measureDepth() int {
	var depth int = 1
	var leftDepth int = 0
	var rightDepth int = 0
	if cursor.left != nil {
		leftDepth = cursor.left.measureDepth()
	}
	if cursor.right != nil {
		rightDepth = cursor.right.measureDepth()
	}
	if leftDepth > rightDepth {
		depth += leftDepth
	} else {
		depth += rightDepth
	}
	return depth
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

func (node *GTreeNode[T]) hasChildren() (bool, bool) {
	var l = (node.left != nil)
	var r = (node.right != nil)
	return l, r
}

// Rebalance Traditionally
func (tree *BinaryTree[T]) Rebalance() {
	// var inOrderList *GList[*GTreeNode[T]]
	var inOrderList *GList[int]
	tree.Traverse(func(nodeValue T, idx int) {
		inOrderList.Push(idx)
	})

	var mid int = (int)(inOrderList.length / 2)
	midItm := inOrderList.GetElemAt(mid)

	newTree := CreateBinaryTree[T]()
	newTree.Insert(tree.FindElemAt(midItm), midItm)
	// wip
}

func (cursor *GTreeNode[T]) Rebalance() {

}

// Adding capabilities of AVL tree
func (tree *BinaryTree[T]) RebalanceAVL() {
	var newRoot = tree.base.RebalanceAVL()
	tree.base = newRoot
}

func (cursor *GTreeNode[T]) RebalanceAVL() *GTreeNode[T] {
	if cursor == nil {
		return nil
	}
	var diff int = cursor.calcDepthDiff()
	if diff > 1 {
		// rotateLeft
		cursor = cursor.rotateLeft()
	} else if diff < -1 {
		// rotateRight
		cursor = cursor.rotateRight()
	}
	var hasLeft, hasRight = cursor.hasChildren()
	if hasLeft {
		cursor.left.RebalanceAVL()
	}
	if hasRight {
		cursor.right.RebalanceAVL()
	}
	return cursor
}

func (root *GTreeNode[T]) rotateLeft() *GTreeNode[T] {
	return nil
}

func (root *GTreeNode[T]) rotateRight() *GTreeNode[T] {
	var t1, t2, t3, t4 *GTreeNode[T]
	// leftNode = root.left
	// rightNode = root.right
	t1 = root.left.left
	t2 = root.left.right.left
	t3 = root.left.right.right
	t4 = root.right

	newRoot := root.left.right
	newRoot.left = root.left
	newRoot.right = root

	newRoot.left.left = t1
	newRoot.left.right = t2
	newRoot.right.left = t3
	newRoot.right.right = t4
	return newRoot
}

func (cursor *GTreeNode[T]) calcDepthDiff() int {
	var leftDepth int = 0
	var rightDepth int = 0
	var hasLeft, hasRight = cursor.hasChildren()
	if hasLeft {
		leftDepth = cursor.measureDepth()
	}
	if hasRight {
		rightDepth = cursor.measureDepth()
	}
	return rightDepth - leftDepth
}

//  this function is incapable of solving for its purpose as of now
func (cursor *GTreeNode[T]) CalcBalance() int {
	// gap := cursor.left.measureDepth() - cursor.right.measureDepth()
	// return gap
	if cursor.depth != 0 {
		return cursor.depth
	}
	var hasLeft, hasRight = cursor.hasChildren()
	var leftDiff = 0
	var rightDiff = 0
	if hasLeft {
		leftDiff = cursor.left.CalcBalance() + 1
	}
	if hasRight {
		rightDiff = cursor.right.CalcBalance() + 1
	}
	fmt.Println("\n Balance at ", cursor.value, leftDiff, rightDiff, "is", rightDiff-leftDiff)
	cursor.depth = rightDiff - leftDiff
	return cursor.depth
}
