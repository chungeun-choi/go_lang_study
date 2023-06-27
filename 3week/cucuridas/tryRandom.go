package cucuridas

import (
	"fmt"
	"math/rand"
	"time"
)


func random() {
	rand.Seed(time.Now().UnixNano())

	n:= rand.Intn(100)

	fmt.Println(n)
}