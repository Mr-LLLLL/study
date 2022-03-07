package bst

import (
	"data_structure/5_chapter/binTree"
	"reflect"
)

type BST struct {
	binTree.BinTree
	Hot *binTree.BinNode
}

// go don't support reference, so here use pointer, in order to effective
func (bst *BST) Search_(e interface{}) **binTree.BinNode {
	n := &bst.Root
	if *n == nil || e == (*n).Data {
		return n
	}
	for true {
		bst.Hot = *n
		if Less(e, (*n).Data) {
			n = &((*n).Lc)
		} else {
			n = &((*n).Rc)
		}
		if *n == nil || e == (*n).Data {
			return n
		}
	}
	return nil
}

func (bst *BST) Search(e interface{}) *binTree.BinNode {
	return *bst.Search_(e)
}

func (bst *BST) Insert(e interface{}) *binTree.BinNode {
	x := bst.Search_(e)
	if *x != nil {
		return *x
	}
	*x = binTree.NewNode(e, bst.Hot)
	bst.Size_++
	bst.Hot.UpdateHeightAbove()
	return *x
}

func (bst *BST) RemoveAt(n *binTree.BinNode) *binTree.BinNode {
	var succ *binTree.BinNode
	dropNode := n
	if n.Lc == nil {
		succ = n.Rc
	} else if n.Rc == nil {
		succ = n.Lc
	} else {
		dropNode = dropNode.Succ()
		dropNode.Data, n.Data = n.Data, dropNode.Data
		p := dropNode.Parent
		if p == n {
			succ = dropNode.Rc
		} else {
			succ = dropNode.Rc
		}
	}
	bst.Hot = dropNode.Parent
	if succ != nil {
		succ.Parent = bst.Hot
	}
	if bst.Hot.Rc == dropNode {
		bst.Hot.Rc = succ
	} else {
		bst.Hot.Lc = succ
	}
	return dropNode
}

func (bst *BST) Remove(e interface{}) *binTree.BinNode {
	node := bst.Search(e)
	if node == nil {
		return nil
	}
	dropNode := bst.RemoveAt(node)
	bst.Size_--
	bst.Hot.UpdateHeightAbove()
	return dropNode
}

func (bst *BST) Size() int {
	return bst.Size_
}

func Less(e1, e2 interface{}) bool {
	if reflect.TypeOf(e1) != reflect.TypeOf(e2) {
		return false
	}
	switch e1 := e1.(type) {
	case int:
		e2 := e2.(int)
		if e1 < e2 {
			return true
		}
	case string:
		e2 := e2.(string)
		if e1 < e2 {
			return true
		}
	default:
		return false
	}
	return false
}

func NewBST() *BST {
	return &BST{
		BinTree: *binTree.NewTree(),
		Hot:     nil,
	}
}
