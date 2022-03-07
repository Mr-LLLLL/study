package splay

import (
	"testing"
)

func Test_Test(t *testing.T) {
	s := NewSplay()
	for i := 1; i < 8; i++ {
		s.Insert(i)
	}
	s.Remove(1)
	s.Remove(1)
	s.Remove(1)
	s.TravLevel()
}
