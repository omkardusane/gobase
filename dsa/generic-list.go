package dsa

/* lists for generic types */
type GNode[T any] struct {
	value T
	next  *GNode[T]
}

type GList[T any] struct {
	first  *GNode[T]
	last   *GNode[T]
	length int32
}

func CreateGList[T any](item T) *GList[T] {
	return &GList[T]{length: 0, first: nil, last: nil}
}

func (list *GList[T]) ForEach(reader func(nodeVal T, idx int)) {
	var idx int = 0
	item := list.first
	for item != nil {
		reader(item.value, idx)
		idx++
		item = item.next
	}
}

func (list *GList[T]) Push(nodeVal T) {
	gNode := &GNode[T]{value: nodeVal, next: nil}
	if list.last == nil {
		list.first = gNode
	} else {
		list.last.next = gNode
	}
	list.last = gNode
	list.length++
}

func (list *GList[T]) Pop() T {
	gNode := list.last
	if gNode != nil {
		return gNode.value
	} else {
		var zeroVal T
		return zeroVal
	}
}
