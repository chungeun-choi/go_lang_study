package cucuridas

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	e4 := v.PushBack(4)

	v.Insertbefore(3, e4)
}