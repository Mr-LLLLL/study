package leetcode

import (
	"reflect"
	"testing"
)

func TestCode_39(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{7}, {2, 2, 3}},
		},
		{
			name: "test2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_39(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Code_39() = %v, want %v", got, tt.want)
			}
		})
	}
}
