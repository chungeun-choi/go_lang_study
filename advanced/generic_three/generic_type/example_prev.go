package generic_type

type PrvNode struct {
	val  interface{}
	next *PrvNode
}

func NewPrevNode(v interface{}) *PrvNode {
	return &PrvNode{
		val: v,
	}
}
