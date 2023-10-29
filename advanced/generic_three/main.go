package main

import (
	"advanced/generic_three/generic_type"
	"fmt"
)

func call1() {
	node1 := generic_type.NewNode(33)
	node1.Push(44).Push(55).Push(66)

	for node1 != nil {
		fmt.Println(node1.Var)
		node1 = node1.Next
	}
}

func main() {
	call1()
}
