package cucuridas

import "fmt"

func add(a,b int) int {
	return a+b
}


func mul(a,b int) int {
	return a*b
}


func geOperator(op string) func (int, int) int {
	if op == "+" {
		return add
	}else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {

	var oprator func (int int) int
	oprator = geOperator("*")

	var result = oprator(3,4)
	fmt.Println(result)
}