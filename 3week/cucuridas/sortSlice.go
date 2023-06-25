package cucuridas

import (
	"fmt"
	"sort"
)


func sortslice() {
	s:= []int{5,2,6,3,1,4}
	sort.Ints(s)
	fmt.Println(s)
}