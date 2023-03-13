package quickSort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				arr: []int{3, 4, 5, 667, 6, 3, 3, 1, 5, 345, 5, 4, 2, 435, 354},
			},
		},
		{
			name: "test2",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
		},
		{
			name: "test3",
			args: args{
				arr: []int{3, 3, 3, 3, 3, 3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			if !sort.IntsAreSorted(tt.args.arr) {
				t.Error("is not sorted")
			}
		})
	}
}
