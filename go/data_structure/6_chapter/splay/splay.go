package splay

import (
	"data_structure/5_chapter/binTree"
	"data_structure/5_chapter/bst"
)

type Splay struct {
	bst.BST
}

func (s *Splay) splay(n *binTree.BinNode) *binTree.BinNode {
	if n == nil {
		return nil
	}
	p := n.Parent
	if p == nil {
		return n
	}
	g := p.Parent
	for n.Parent != nil && n.Parent.Parent != nil {
		p = n.Parent
		g = p.Parent
		gg := g.Parent // great-grand parent
		if p.IsLChild() {
			if n.IsLChild() { // zig-zig
				attachAsLChild(g, p.Rc)
				attachAsLChild(p, n.Rc)
				attachAsRChild(p, g)
				attachAsRChild(n, p)
			} else { // zag-zig
				attachAsRChild(p, n.Lc)
				attachAsLChild(g, n.Rc)
				attachAsLChild(n, p)
				attachAsRChild(n, g)
			}
		} else {
			if n.IsLChild() { // zig-zag
				attachAsRChild(g, n.Lc)
				attachAsLChild(p, n.Rc)
				attachAsLChild(n, g)
				attachAsRChild(n, p)
			} else { // zag-zag
				attachAsRChild(g, p.Lc)
				attachAsRChild(p, n.Lc)
				attachAsLChild(p, g)
				attachAsLChild(n, p)
			}
		}
		if gg != nil {
			if g == gg.Lc {
				attachAsLChild(gg, n)
			} else {
				attachAsRChild(gg, n)
			}
		} else {
			n.Parent = nil
		}
		g.UpdateHeight()
		p.UpdateHeight()
		n.UpdateHeight()
	}
	p = n.Parent
	if p != nil {
		if n.IsLChild() {
			attachAsLChild(p, n.Rc)
			attachAsRChild(n, p)
		} else {
			attachAsRChild(p, n.Lc)
			attachAsLChild(n, p)
		}
		p.UpdateHeight()
		n.UpdateHeight()
	}
	n.Parent = nil
	return n
}

func (s *Splay) Search(e interface{}) *binTree.BinNode {
	n := s.Search_(e)
	if *n == nil {
		n = &s.Hot
	}
	s.Root = s.splay(*n)
	return s.Root
}

func (s *Splay) Insert(e interface{}) *binTree.BinNode {
	if s.Root == nil {
		s.Size_++
		s.Root = binTree.NewNode(e)
	}
	if e == s.Search(e).Data {
		return s.Root
	}
	s.Size_++
	t := s.Root
	if bst.Less(s.Root.Data, e) {
		s.Root = binTree.NewNode(e, nil, t, t.Rc)
		t.Parent = s.Root
		if t.Rc != nil {
			t.Rc.Parent = s.Root
			t.Rc = nil
		}
	} else {
		s.Root = binTree.NewNode(e, nil, t.Lc, t)
		t.Parent = s.Root
		if t.Lc != nil {
			t.Lc.Parent = s.Root
			t.Lc = nil
		}
	}
	t.UpdateHeightAbove()
	return s.Root
}

func (s *Splay) Remove(e interface{}) *binTree.BinNode {
	if s.Root == nil {
		return nil
	}
	if e != s.Search(e).Data {
		return s.Root
	}
	w := s.Root
	if s.Root.Lc == nil {
		s.Root = s.Root.Rc
		if s.Root != nil {
			s.Root.Parent = nil
		}
	} else if s.Root.Rc == nil {
		s.Root = s.Root.Lc
		if s.Root != nil {
			s.Root.Parent = nil
		}
	} else {
		lTree := s.Root.Lc
		lTree.Parent = nil
		s.Root.Lc = nil
		s.Root = s.Root.Rc
		s.Root.Parent = nil
		s.Search(e)
		s.Root.Lc = lTree
		lTree.Parent = s.Root
	}
	s.Size_--
	if s.Root != nil {
		s.Root.UpdateHeight()
	}
	return w
}

func attachAsLChild(p, lc *binTree.BinNode) {
	p.Lc = lc
	if lc != nil {
		lc.Parent = p
	}
}

func attachAsRChild(p, rc *binTree.BinNode) {
	p.Rc = rc
	if rc != nil {
		rc.Parent = p
	}
}

func NewSplay() *Splay {
	return &Splay{
		BST: *bst.NewBST(),
	}
}
