package binTree

import (
	"testing"
)

func Test_BinTree(t *testing.T) {
	tree := NewTree()
	for i := 0; i < 20; i++ {
		tree.insert(i)
	}
	tree.TravPost()
}
