package avl

import (
	"data_structure/5_chapter/binTree"
	"data_structure/5_chapter/bst"
	"math"
)

type AVL struct {
	bst.BST
}

func (a *AVL) Insert(e interface{}) *binTree.BinNode {
	x := a.Search_(e)
	if *x != nil {
		return *x
	}
	*x = binTree.NewNode(e, a.Hot)
	a.Size_++
	for node := a.Hot; node != nil; node = node.Parent {
		if !avlBalanced(node) {
			p := node.Parent
			if p == nil {
				a.Root = rorateAt(node.TallerChild().TallerChild())
			} else {
				if p.Lc == node {
					p.Lc = rorateAt(node.TallerChild().TallerChild())
				} else {
					p.Rc = rorateAt(node.TallerChild().TallerChild())
				}
			}
			break
		} else {
			node.UpdateHeight()
		}
	}
	return *x
}

func (a *AVL) Remove(e interface{}) *binTree.BinNode {
	x := a.Search(e)
	if x == nil {
		return nil
	}
	dropNode := a.RemoveAt(x)
	for n := a.Hot; n != nil; n = n.Parent {
		if !avlBalanced(n) {
			p := n.Parent
			if p == nil {
				a.Root = rorateAt(n.TallerChild().TallerChild())
			} else {
				if p.Lc == n {
					p.Lc = rorateAt(n.TallerChild().TallerChild())
				} else {
					p.Rc = rorateAt(n.TallerChild().TallerChild())
				}
			}
		}
		n.UpdateHeight()
	}
	return dropNode
}

func avlBalanced(node *binTree.BinNode) bool {
	return math.Abs(float64(binTree.Stature(node.Lc)-binTree.Stature(node.Rc))) < 2
}

func rorateAt(n *binTree.BinNode) *binTree.BinNode {
	p, g := n.Parent, n.Parent.Parent
	if p.IsLChild() {
		if n.IsLChild() {
			p.Parent = g.Parent
			return connect34(p, n, g, n.Lc, n.Rc, p.Rc, g.Rc)
		} else {
			n.Parent = g.Parent
			return connect34(n, p, g, p.Lc, n.Lc, n.Rc, g.Rc)
		}
	} else {
		if n.IsLChild() {
			n.Parent = g.Parent
			return connect34(n, g, p, g.Lc, n.Lc, n.Rc, p.Rc)
		} else {
			p.Parent = g.Parent
			return connect34(p, g, n, g.Lc, p.Lc, n.Lc, n.Rc)
		}
	}
}

func connect34(a, b, c, t0, t1, t2, t3 *binTree.BinNode) *binTree.BinNode {
	b.Lc = t0
	if t0 != nil {
		t0.Parent = b
	}
	b.Rc = t1
	if t1 != nil {
		t1.Parent = b
	}
	b.UpdateHeight()

	c.Lc = t2
	if t2 != nil {
		t2.Parent = c
	}
	c.Rc = t3
	if t3 != nil {
		t3.Parent = c
	}
	c.UpdateHeight()

	a.Lc = b
	b.Parent = a
	a.Rc = c
	c.Parent = a
	a.UpdateHeight()

	return a
}

func NewAVL() *AVL {
	return &AVL{
		BST: *bst.NewBST(),
	}
}
