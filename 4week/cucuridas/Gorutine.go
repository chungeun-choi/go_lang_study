package cucuridas

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintHangul() {
	hanguls := []rune {'가','나','다','라','마'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c",v)
	}
}

func PrintNumbers() {
	for i:=1; i <=5;i++{
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d",i)
	}
}


func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3* time.Second)
}