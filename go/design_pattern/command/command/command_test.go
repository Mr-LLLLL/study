package command

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	a := []int{1, 2}
	a = append(a[:1], a[1+1:]...)
	a = a[:0]
	fmt.Println(a)
}
