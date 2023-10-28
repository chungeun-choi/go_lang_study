package main

import (
	c1 "advanced/generic_two/check_interface_limitedvalue"
	c3 "advanced/generic_two/example_map"
	"advanced/generic_two/limited_interface"
	"fmt"
)

type MyString struct {
	name string
}

func (s MyString) String() string {
	return s.name
}

func Call1() {
	m := MyString{"test"}
	c1.Show2(m)

}

func Call2() {
	var intObject limited_interface.IntStruct = 33

	result := intObject.Sum(123)
	fmt.Println(result)
}

func Call3() {
	IntList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := c3.Map(IntList, func(v int) int {
		return v * 2
	})

	fmt.Println(result)
}

func main() {
	Call1()
	Call2()
	Call3()
}
