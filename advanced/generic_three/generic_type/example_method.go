package generic_type

type NodeMe struct {
	Val  int
	Next *NodeMe
}

func NewNodeMe(v int) *NodeMe {
	return &NodeMe{
		Val: 22,
	}
}

/*
// 타입 매개변수 단일로는 메소드에 사용할 수 없음!
func (n *NodeMe) Push[T any](v T) {

}
*/
