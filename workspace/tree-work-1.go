package workspace

import (
	"fmt"
	"gobase_om/dsa"
)

func ExecPlayList4() {
	fmt.Println("[Trees]")
	tr1 := dsa.CreateBinaryTree[string]()
	tr1.Insert("A", 21)
	tr1.Insert("B", 23)
	tr1.Insert("C", 20)
	tr1.Insert("D", 24)
	tr1.Insert("E", 17)
	tr1.Insert("F", 18)
	tr1.Insert("G", 22)
	counter := 0
	tr1.Traverse(func(nodeValue string, idx int) {
		counter++
		fmt.Printf("\n %d: { index: %d, val : %s}, ", counter, idx, nodeValue)
	})
	fmt.Println("\n[BFS Traverse]")
	counter = 0
	tr1.BfsTraverse(func(nodeValue string, depth int) {
		counter++
		fmt.Printf("\n %d: { depth: %d, val : %s}, ", counter, depth, nodeValue)
	})
	fmt.Print("\n\n")
	treeBalance := tr1.GetRoot().CalcBalance()
	fmt.Println("\n Balance in tree", treeBalance)

	// treeBalance := tr1.GetRoot().Left().CalcBalance()
	// fmt.Println("\n Balance in tree", treeBalance)

	// treeBalance := tr1.GetRoot().CalcBalance()
	// fmt.Println("\n Balance in tree", treeBalance)
}
