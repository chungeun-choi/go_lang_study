package one

import (
	"advanced/module_package/two"
	"fmt"
)

func PrintOnePackage() {
	two.PrintTwoPackage()
	fmt.Println("It is text by using PrintOnePakcage func")
}
