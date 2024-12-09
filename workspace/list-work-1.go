package workspace

import (
	"fmt"
	. "gobase_om/dsa"
	. "gobase_om/greetings"
	"gobase_om/t1"
)

func Exec() {
	testHelloYellow()
	playlist1()
	genList()
}

func testHelloYellow() {
	var a4 = Hello("12")
	var dif = a4*9 - t1.Yellow()*9
	fmt.Println("DIfF ")
	fmt.Println(a4, a4*9, dif/9)
}

func playlist1() {
	fmt.Println("initializing list")

	i1 := CreateNewItem(101)
	i1 = i1.AddNext(122)
	i1 = i1.AddNext(123)
	i1 = i1.AddNext(124)

	i3 := i1.AddNext(131)
	i3.AddNext(137)

	fmt.Println("Printing list")
	i1.ForEach(func(val int, idx int) {
		fmt.Printf("%d : %d \n", idx, val)
	})
}

func genList() {

	var goOn bool = true
	var newNum int
	l2Starter := CreateNewItem(301)
	var l2cursor *Item = l2Starter
	// Prompt the user for their name
	for goOn {
		fmt.Print("Enter next number: ")
		fmt.Scanln(&newNum) // Read the input into the 'name' variable
		l2cursor.AddNext(newNum)
		fmt.Print("Add one more? Y/N (default Y): ")
		var tempAns string
		fmt.Scanln(&tempAns)
		if tempAns == "N" || tempAns == "n" {
			goOn = false
			// break;
		} else {
			l2cursor = l2cursor.GetNext()
			goOn = true
		}
	}
	fmt.Println("We are done. Printing the list again")
	l2Starter.PrintAll()
}
