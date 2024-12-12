package workspace

import (
	"fmt"
	"gobase_om/dsa"
)

func ExecPlayList4() {
	fmt.Println("[Trees]")
	tr1 := dsa.CreateBinaryTree[int]()
	tr1.Insert(2100, 21)
	tr1.Insert(2278, 22)
	tr1.Insert(2020, 20)
	tr1.Insert(3001, 23)
	counter := 0
	tr1.Traverse(func(nodeValue int, idx int) {
		counter++
		fmt.Printf("\n %d: { index: %d, val : %d}, ", counter, idx, nodeValue)
	})
}
