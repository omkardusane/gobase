package dsa

type Node struct {
	element Element
	left    *Node
	right   *Node
}

type Element struct {
	value int
}
