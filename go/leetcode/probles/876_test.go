package probles

import (
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
			},
		},
		{
			name: "test2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleNode(tt.args.head)
			got1 := got
			should := tt.want
			should1 := should
			suc := true
			for got != nil {
				if should == nil || got.Val != should.Val {
					suc = false
					break
				}
				got = got.Next
				should = should.Next
			}
			if !suc {
				gotA := make([]int, 0)
				for got1 != nil {
					gotA = append(gotA, got1.Val)
					got1 = got1.Next
				}
				shoudA := make([]int, 0)
				for should1 != nil {
					shoudA = append(shoudA, should1.Val)
					should1 = should1.Next
				}

				t.Errorf("Code_876() = %v, want %v", gotA, shoudA)
			}
		})
	}
}
