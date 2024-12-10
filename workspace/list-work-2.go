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
	fmt.Printf("really good sum = %d", sum)
}

func ExecPlayList3() {

}
