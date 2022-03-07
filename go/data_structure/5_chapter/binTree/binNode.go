package binTree

import (
	"data_structure/5_chapter/queue"
	"data_structure/5_chapter/stack"
	"fmt"
	"log"
)

type color int

const (
	RB_RED = iota
	RB_BLACK
)

type BinNode struct {
	Data           interface{}
	Parent, Lc, Rc *BinNode
	Height         int
	Color          color
}

func (node *BinNode) insertAsLC(data interface{}) *BinNode {
	node.Lc = NewNode(data, node)
	return node.Lc
}

func (node *BinNode) insertAsRC(data interface{}) *BinNode {
	node.Rc = NewNode(data, node)
	return node.Rc
}

func (node *BinNode) remove() (data interface{}) {
	if node == nil {
		return nil
	}
	if node.Parent.Rc == node {
		node.Parent.Rc = nil
	} else {
		node.Parent.Lc = nil
	}
	data = node.Data
	node = nil
	return
}

func (node *BinNode) Succ() *BinNode {
	if node.Rc != nil {
		node = node.Rc
		for node.Lc != nil {
			node = node.Lc
		}
	} else {
		for node.Parent != nil && node.Parent.Rc == node {
			node = node.Parent
		}
		node = node.Parent
	}
	return node
}

func (node *BinNode) isLeaf() bool {
	if node.Lc == nil || node.Rc == nil {
		return true
	} else {
		return false
	}
}

func (node *BinNode) travPre_I() {
	s := stack.New()
	for true {
		for node != nil {
			visitNode(node)
			s.Push(node.Rc)
			node = node.Lc
		}
		if s.Empty() {
			break
		}
		tmp, err := s.Pop()
		node = tmp.(*BinNode)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (node *BinNode) travPre_R() {
	if node == nil {
		return
	}
	visitNode(node)
	node.Lc.travPre_R()
	node.Rc.travPre_R()
}

func (node *BinNode) travIn_I() {
	s := stack.New()
	for true {
		for node != nil {
			s.Push(node)
			node = node.Lc
		}
		if s.Empty() {
			break
		}
		tmp, err := s.Pop()
		if err != nil {
			return
		}
		node = tmp.(*BinNode)
		visitNode(node)
		node = node.Rc
	}
}

func (node *BinNode) travIn_R() {
	if node == nil {
		return
	}
	node.Lc.travIn_R()
	visitNode(node)
	node.Rc.travIn_R()
}

func (node *BinNode) travPost_I() {
	s := stack.New()
	if node != nil {
		s.Push(node)
	} else {
		return
	}
	for !s.Empty() {
		tmp, err := s.Top()
		if err != nil {
			log.Fatal(err)
		}
		topNode := tmp.(*BinNode)
		// if top don't equal node's parent, top is right child
		if topNode != node.Parent {
			for true {
				if topNode.Lc == nil && topNode.Rc == nil {
					break
				}
				if topNode.Rc != nil {
					s.Push(topNode.Rc)
				}
				if topNode.Lc != nil {
					s.Push(topNode.Lc)
				}
				tmp, err := s.Top()
				if err != nil {
					log.Fatal(err)
				}
				topNode = tmp.(*BinNode)
			}
		}
		tmp, err = s.Pop()
		node = tmp.(*BinNode)
		visitNode(node)
	}
}

func (node *BinNode) travPost_R() {
	if node == nil {
		return
	}
	node.Lc.travPost_R()
	node.Rc.travPost_R()
	visitNode(node)
}

func (node *BinNode) travLevel() {
	q := queue.New()
	if node != nil {
		q.Enqueue(node)
	} else {
		fmt.Println("node is nil")
	}
	for q.Size() != 0 {
		n, err := q.Dequeue()
		if err != nil {
			log.Fatal(err)
		}
		n1 := n.(*BinNode)
		if n1.Lc != nil {
			q.Enqueue(n1.Lc)
		}
		if n1.Rc != nil {
			q.Enqueue(n1.Rc)
		}
		visitNode(n.(*BinNode))
	}
}

func (node *BinNode) UpdateHeight() int {
	if node == nil {
		return -1
	}
	oldHeight := Stature(node)
	if Stature(node.Lc) > Stature(node.Rc) {
		node.Height = Stature(node.Lc) + 1
	} else {
		node.Height = Stature(node.Rc) + 1
	}
	return oldHeight
}

func (node *BinNode) UpdateHeightAbove() {
	n := node
	for n != nil {
		n.UpdateHeight()
		n = n.Parent
	}
}

func (n *BinNode) IsLChild() bool {
	if n.Parent == nil {
		return false
	}
	return n.Parent.Lc == n
}

func (n *BinNode) IsRChild() bool {
	if n.Parent == nil {
		return false
	}
	return n.Parent.Rc == n
}

func (n *BinNode) TallerChild() *BinNode {
	if Stature(n.Lc) > Stature(n.Rc) {
		return n.Lc
	} else if Stature(n.Rc) > Stature(n.Lc) {
		return n.Rc
	} else {
		if n.IsLChild() {
			return n.Lc
		} else {
			return n.Rc
		}
	}
}

func visitNode(n *BinNode) {
	fmt.Printf("%d, %d\n", n.Data, n.Height)
}

func Stature(n *BinNode) int {
	if n == nil {
		return -1
	} else {
		return n.Height
	}
}

func NewNode(data interface{}, args ...*BinNode) *BinNode {
	nodes := make([]*BinNode, 3)
	copy(nodes, args)
	return &BinNode{
		Data:   data,
		Parent: nodes[0],
		Lc:     nodes[1],
		Rc:     nodes[2],
		Height: 0,
		Color:  RB_RED,
	}
}
