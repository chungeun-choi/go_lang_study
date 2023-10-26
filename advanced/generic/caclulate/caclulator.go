package caclulate

import "golang.org/x/exp/constraints"

/*
아래의 두개의 함수는 강 타입 언어인 go 언어의 특징으로인해 각각 다른 인자를 받아 처리하도록
따로 구현을 해야함을 보여줌
*/
func SumInt(a, b int) int {
	value := a + b
	return value
}

func SumFload(a, b float32) float32 {
	value := a + b
	return value
}

/*
// 해당 함수는 interface로 인자 값을 받음으로 인해 내부 연산 로직에 사용되는 '+'를 사용할 수 없음
func SumInterface(a,b interface{}) interface {
	value  := a + b
	return value
}
*/

/*
Generic 프로그래밍 사용방법
*/
func SumExample[T any](a, b T) (T, T) {
	value := a
	value2 := b
	return value, value2
}

/*
//여전히 계산해 필요한 값은 any로 되어 있기에 사용할 수 없음
func Sum[T any](a, b T) T {
	value := a+b
	return value
}
*/

// 위 문제 해결 방법
func Sum[T int | float32 | float64 | int16](a, b T) T {
	value := a + b
	return value
}

// 타입 제한자 미리 정의해 두기
type T interface {
	int8 | int16 | int32 | float32 | float64
}

func SumLimit[tttt T](a, b tttt) tttt {
	value := a + b
	return value
}

// go언어는 기본적인 타입 제한자를 제공해줌
func SumConstraints[T constraints.Ordered](a, b T) T {
	value := a + b
	return value
}
