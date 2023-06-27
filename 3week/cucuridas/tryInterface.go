package cucuridas

import (
	"fmt"

	v1 "k8s.io/client-go/informers/batch/v1"
)

type CarProduct interface {
	getName() string
}


type Genesis struct {
	Name string
}

func (g *Genesis) getName() string {
	return g.Name
}

func (g *Genesis) setName(name string) {
	g.Name = name
}

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n",nums)
	for _,v:= range nums {
		sum +=v
	}

	return sum
}

func main() {
	genesis := &Genesis{ Name: "Mycar"}

	var product CarProduct

	product = genesis

	fmt.Printf("%s\n",product.getName())
	
}