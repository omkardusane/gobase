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
	printTree(tr1)
	fmt.Println("\n[BFS Traverse]")
	counter := 0
	tr1.BfsTraverse(func(nodeValue string, depth int) {
		counter++
		fmt.Printf("\n %d: { depth: %d, val : %s}, ", counter, depth, nodeValue)
	})
	// treeBalance := tr1.GetRoot().CalcBalance()
	// fmt.Println("\n Balance in tree", treeBalance)
	fmt.Print("\n")

	tr1.RebalanceAVL()
	printTree(tr1)
}

func printTree(tree *dsa.BinaryTree[string]) {
	counter := 0
	tree.Traverse(func(nodeValue string, idx int) {
		counter++
		fmt.Printf("\n %d: { val : %s}, ", counter, nodeValue)
	})
}
