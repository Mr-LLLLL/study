package btree

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	s1 := []*int{}
	s := s1[:0]
	s1 = append(s1, s[0:]...)
	fmt.Println(len(s), len(s1))
}
