package probles

import (
	"fmt"
	"testing"
)

func TestMinStack_MinStack(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := Constructor()
			stack.Push(-2)
			stack.Push(0)
			stack.Push(1)
			stack.Push(-3)
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Top())
			fmt.Println(stack.Min())
		})
	}
}
