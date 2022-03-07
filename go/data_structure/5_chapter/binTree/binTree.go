package binTree

import (
	"math/rand"
	"time"
)

type BinTree struct {
	Size_ int
	Root  *BinNode
}

func (t *BinTree) Insert(data interface{}) {
	t.Size_++
	n := t.getRandomLeaf()
	if n == nil {
		t.Root = NewNode(data)
		return
	}
	if rand.Intn(2) == 0 {
		if n.Lc == nil {
			n.insertAsLC(data)
		} else {
			n.insertAsRC(data)
		}
	} else {
		if n.Rc == nil {
			n.insertAsRC(data)
		} else {
			n.insertAsLC(data)
		}
	}
}

func (t *BinTree) Remove(node *BinNode) interface{} {
	return node.remove()
}

func (t *BinTree) getRandomLeaf() *BinNode {
	rand.Seed(time.Now().Unix())
	n := t.Root
	if n == nil {
		return nil
	}
	for !n.isLeaf() {
		if rand.Intn(2) == 0 {
			n = n.Lc
		} else {
			n = n.Rc
		}
	}
	return n
}

func (t *BinTree) TravPost() {
	r := t.Root
	// r.travPost_R()
	r.travPost_I()
}

func (t *BinTree) TravIn() {
	r := t.Root
	// r.travIn_R()
	r.travIn_I()
}

func (t *BinTree) TravPre() {
	r := t.Root
	// r.travPre_R()
	r.travPre_I()
}

func (t *BinTree) TravLevel() {
	r := t.Root
	r.travLevel()
}

func NewTree() *BinTree {
	return &BinTree{
		Root:  nil,
		Size_: 0,
	}
}
