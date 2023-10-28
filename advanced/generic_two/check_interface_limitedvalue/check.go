package check_interface_limitedvalue

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// String 함수를 구현하도록 제한하는 인터페이스
type ExampleInt interface {
	String() string
}

// int와 같은 숫자 계산을 할 수 있는 값으로 제한하는 타입 제한자
type ExampleLimit interface {
	constraints.Ordered
}

/*
'ExampleInt' 인터페이스를 타입 제한자로 둠으로써 String() 함수가'
구현되어진 모든 객체들은 Show 함수의 인자 값으로 입력이 가능
하지만 기존 인터페이스 사용하는 것이 목적에 더 알맞다
*/
func Show1[T ExampleInt](b T) {
	fmt.Println(b.String())
}

/*
아래의 함수는 타입제한자를 타입(인터페이스)로 받기 때문에 사용할 수 없음
*/
func Show2(b ExampleInt) {
	fmt.Println(b)
}
