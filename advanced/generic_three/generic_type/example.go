package generic_type

// 구조체 생성 시 generic 타입 지정 변수를 통해 val 필드의 타입을 자유롭게 지정하도록 설정
type Node[T any] struct {
	Var  T
	Next *Node[T]
}

// 생성자 함수를 통해 입력 받는 파라미터(param)의 타입을 전달 받아 Node 객체 생성할 때 타입 지정 변수로 사용
func NewNode[T any](param T) *Node[T] {
	return &Node[T]{
		Var: param,
	}
}

/*
		타입 지정 변수는 해당 객체 내에서 지속적으로 사용됨
	 	해당 타입 지정변수는 메소드에 따로 지정된 것이 아닌 구조체를 통해 객체를 생성할 떄
		정의되어진 타입 지정변수를 사용함
		메소드만 따로 타입 지정변수를 정의하여 사용할 수 없음
*/
func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.Next = node
	return node
}
