package avl

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	avl := NewAVL()
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(2)
	avl.Insert(8)
	avl.Insert(7)
	avl.Insert(4)
	avl.Insert(9)
	avl.TravLevel()
	fmt.Println()
	avl.Remove(7)
	avl.Remove(3)
	avl.Remove(2)
	avl.TravLevel()
}

//             3
//     2				7
// 1				4		8
//                             9
