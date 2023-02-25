package probles

import (
	"fmt"
	"testing"
)

func TestCode_1339(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				root: getRootNode([]int{1, 2, 3, 4, 5, 6}),
			},
			want: 110,
		},
		{
			name: "test2",
			args: args{
				root: getRootNode([]int{1, 0, 2, 3, 4, 0, 0, 5, 6}),
			},
			want: 90,
		},
		{
			name: "test3",
			args: args{
				root: getRootNode([]int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1}),
			},
			want: 1025,
		},
		{
			name: "test4",
			args: args{
				root: getRootNode([]int{1, 1}),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1339(tt.args.root); got != tt.want {
				t.Errorf("Code_1339() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getRootNode(in []int) *TreeNode {
	root := &TreeNode{
		Val: in[0],
	}

	queue := make([]*TreeNode, 0, len(in))
	queue = append(queue, root)
	j := 1
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if getArrVAl(in, j) != 0 {
			node.Left = &TreeNode{
				Val: getArrVAl(in, j),
			}
			queue = append(queue, node.Left)
		} else {
			node.Left = nil
		}
		j++

		if getArrVAl(in, j) != 0 {
			node.Right = &TreeNode{
				Val: getArrVAl(in, j),
			}
			queue = append(queue, node.Right)
		} else {
			node.Right = nil
		}
		j++
	}

	return root
}

func getArrVAl(in []int, index int) int {
	if index >= len(in) {
		return 0
	}

	return in[index]
}

func printnode(root *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for i := 0; i < len(queue); i++ {
		fmt.Println(queue[i].Val)
		node := queue[i]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
