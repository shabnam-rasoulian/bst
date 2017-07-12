package main

import ( 
	"fmt"
	"github.com/shabnam-rasoulian/bst"
)

type ComparableInt int

func (i ComparableInt) Compare(that interface{}) int {
	ti, ok := that.(ComparableInt)
	if !ok {
		panic("not an int")
	}
	if ti == i {
		return 0
	} else if ti < i {
		return -1
	} else {
		return 1
	}
}

func main() {
	t := bst.Tree{}
	t.Insert(ComparableInt(50))
	t.Insert(ComparableInt(20))
	t.Insert(ComparableInt(70))
	t.Insert(ComparableInt(100))
	fmt.Println(t.String())
}
