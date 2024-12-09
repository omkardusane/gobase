package workspace

import (
	"fmt"
	"gobase_om/dsa"
)

func ExecPlayList2() {
	var l1 *dsa.List = dsa.CreateNewList()
	l1.AddNext(11)
	l1.AddNext(12)
	l1.AddNext(13)
	l1.AddNext(14)
	l1.ForEach(func(item *dsa.Item, idx int) {
		fmt.Println("-- ", idx, item.GetValue())
	})
}
