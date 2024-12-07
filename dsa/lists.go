package dsa

import "fmt"

type Item struct {
	value int
	next  *Item
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

func (node *Item) AddNext(val int) *Item {
	node.next = CreateNewItem(val)
	return node.next
}

func (node *Item) GetNext() *Item {
	return node.next
}

func CreateNewItem(val int) *Item {
	// panic("unimplemented")
	var a = Item{value: val | 0, next: nil}
	return &a
}
