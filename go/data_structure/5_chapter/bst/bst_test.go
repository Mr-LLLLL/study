package bst

import (
	"testing"
)

func Benchmark_Test(b *testing.B) {
	bst := NewBST()
	bst.Insert(3)
	bst.Insert(31)
	bst.Insert(43)
	bst.Insert(33)
	bst.Insert(43)
	bst.Insert(83)
	bst.Insert(23)
	bst.Insert(73)
	bst.Insert(22)
	for i := 0; i < b.N; i++ {
		// bst.TravLevel()
		// bst.TravPre()
		bst.TravPost()
		// bst.TravIn()
	}
}

func Test_Test(t *testing.T) {
	bst := NewBST()
	// bst.Insert(3)
	// bst.Insert(31)
	// bst.Insert(43)
	// bst.Insert(33)
	// bst.Insert(43)
	// bst.Insert(83)
	// bst.Insert(23)
	// bst.Insert(73)
	// bst.Insert(22)
	for i := 0; i < 20; i++ {
		bst.Insert(i)
	}

	bst.Remove(2)
	bst.Remove(3)
	// bst.TravLevel()
	// bst.TravPre()
	// bst.TravPost()
	bst.TravIn()
}
