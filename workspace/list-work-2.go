package workspace

import (
	"fmt"
	"gobase_om/dsa"
	"math/rand/v2"
)

func ExecPlayList2() {
	var l1 *dsa.List = dsa.CreateNewList()
	l1.AddNext(11)
	l1.AddNext(12)
	l1.AddNext(int(rand.Int32N(100)))
	l1.AddNext(int(rand.Int32N(100)))
	l1.AddNext(int(rand.Int32N(100)))
	l1.AddNext(int(rand.Int32N(100)))
	l1.ForEach(func(item *dsa.Item, idx int) {
		fmt.Println("-- ", idx, item.GetValue())
	})

	/* The usage of map function */
	var sum = 0
	l1.Map(func(item *dsa.Item, idx int) {
		// item
		sum += item.GetValue()
	})
	fmt.Printf("really good sum = %d \n", sum)
}

func ExecPlayList3() {
	fmt.Println("[ExecPlayList3]")
	l1 := dsa.CreateGList[int]()
	l1.Push(22)
	l1.Push(23)
	//fmt.Println("size = ", l1.GetLength())
	l1.Push(25)
	l1.Push(int(rand.Int32N(100)))
	l1.Push(309)
	fmt.Println("size = ", l1.GetLength())
	l1.ForEach(func(item int, idx int) {
		fmt.Print("( ", item, " @", idx, ")")
		if l1.GetLength() == idx+1 {
			fmt.Print(".")
		} else {
			fmt.Print("->")
		}
	})
	popped := l1.Pop()
	fmt.Println("\n Just Popped ", popped)
	fmt.Println("size = ", l1.GetLength())
	l1.ForEach(func(item int, idx int) {
		fmt.Print("( ", item, " @", idx, ")")
		if l1.GetLength() == idx+1 {
			fmt.Print(".")
		} else {
			fmt.Print("->")
		}
	})

	fmt.Print("\n List of strings \n")
	l2 := dsa.CreateGList[string]()
	l2.Push("Some")
	l2.Push("working")
	l2.Push("lists")
	l2.Push("maybe")
	l2.Pop()
	l2.ForEach(func(item string, idx int) {
		fmt.Print(item)
		if l1.GetLength() == idx+1 {
			fmt.Print(".")
		} else {
			fmt.Print(" ")
		}
	})
}
