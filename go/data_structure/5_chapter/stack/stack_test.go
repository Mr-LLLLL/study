package stack

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	s := New()
	fmt.Println(s.Push(1))
	fmt.Println(s.Push(2))
	fmt.Println(s.Push(3))
	fmt.Println(s.Push(4))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	return
}
