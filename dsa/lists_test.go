package dsa

import "testing"

func TestCreation(t *testing.T) {
	var l1 = CreateNewItem(21)
	if l1.GetValue() == 21 {

	} else {
		t.Fatal("Expected to match the value")
	}
	t.Log("TestCreation passed successfully!")
	count := 0
	l1.ForEach(func(itemVal int, idx int) {
		count++
	})
	if count != 1 {
		t.Fatal("count is incorrect")
	}
}

func TestAdd(t *testing.T) {
	var l1 = CreateNewItem(21)
	var last = l1.AddNext(22).AddNext(23)

	if last.value != 23 {
		t.Fatal("last value does not match")
	}
	t.Log("TestAdd passed successfully!")
	l1.PrintAll()
}
