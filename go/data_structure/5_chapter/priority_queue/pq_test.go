package pq

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	q := NewPq()
	for i := 0; i < 20; i++ {
		q.Insert(i)
	}
	for !q.Empty() {
		fmt.Println(q.DelMax())
	}
}
