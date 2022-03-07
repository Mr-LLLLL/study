package btree

type bNode struct {
	key    []int
	parent *bNode
	child  []*bNode
}

func (n *bNode) find(e int) int {
	r := len(n.key)
	for i, v := range n.key {
		if v >= e {
			r = i
			break
		}
	}
	return r
}

func newNode() (res *bNode) {
	return &bNode{
		key:    []int{},
		parent: nil,
		child:  []*bNode{nil},
	}
}

func NewNode(e int, args ...*bNode) (res *bNode) {
	nodes := make([]*bNode, 3) // default value is nil
	copy(nodes, args)
	res = &bNode{
		key:    []int{e},
		parent: nodes[0],
		child:  nodes[1:3],
	}
	if res.child[1] != nil {
		res.child[1].parent = res
	}
	if res.child[2] != nil {
		res.child[2].parent = res
	}
	return
}
