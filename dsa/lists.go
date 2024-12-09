package dsa

import "fmt"

type Item struct {
	value int
	next  *Item
}

type List struct {
	start  *Item
	last   *Item
	length int
}

func CreateNewList() *List {
	return &List{start: nil, last: nil, length: 0}
}

func (list *List) AddNext(val int) {
	item := CreateNewItem(val)
	if list.last == nil {
		list.start = item
	} else {
		list.last.next = item
	}
	list.last = item
	list.length++
}

func (list *List) ForEach(reader func(item *Item, idx int)) {
	var idx int = 0
	item := list.start
	for item != nil {
		reader(item, idx)
		idx++
		item = item.next
	}
}

func (list *List) Map(reader func(item *Item, idx int)) {
	var idx int = 0
	item := list.start
	for item != nil {
		reader(item, idx)
		idx++
		item = item.next
	}
}

func (list *List) GetCursor() *Item {
	return list.start
}

func CreateNewItem(val int) *Item {
	// panic("unimplemented")
	var a = Item{value: val | 0, next: nil}
	return &a
}

func (base *Item) ForEach(reader func(itemVal int, idx int)) {
	// base.next
	fmt.Println("Traversing list ---")
	var idx int = 0
	item := base
	for item != nil {
		reader(item.value, idx)
		idx++
		item = item.next
	}
}

func (base *Item) PrintAll() {
	fmt.Println("\nTraversing list ---")
	var idx int = 0
	item := base
	for item != nil {
		fmt.Print(idx, "->", item.value, ", ")
		idx++
		item = item.next
	}
	fmt.Println("Ended list ---")
}

/* Methods for each node */
func (node *Item) AddNext(val int) *Item {
	node.next = CreateNewItem(val)
	return node.next
}

func (node *Item) GetNext() *Item {
	return node.next
}

func (node *Item) GetValue() int {
	return node.value
}
