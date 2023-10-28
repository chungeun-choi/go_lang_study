package limited_interface

import "golang.org/x/exp/constraints"

// interface에 타입 제한자를 둠으로써 제한
type ParamOrdered interface {
	constraints.Ordered
	Sum(b int) int
}

// 실제 인터페이스 구현체
type IntStruct int

func (s IntStruct) Sum(b int) int {
	return int(s) + b
}
