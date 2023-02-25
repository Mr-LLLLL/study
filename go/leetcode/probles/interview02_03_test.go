package probles

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name   string
		args   args
		wanted []int
	}{
		{
			name: "test1",
			args: args{
				node: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
		{
			name: "test2",
			args: args{
				node: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 8,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			deleteNode(tt.args.node)
			for tt.args.node != nil {
				fmt.Println(tt.args.node.Val)
				tt.args.node = tt.args.node.Next
			}
		})
	}
}
