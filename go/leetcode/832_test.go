package main

import (
	"reflect"
	"testing"
)

func TestCode_832(t *testing.T) {
	type args struct {
		image [][]int
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
				image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			},
			want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			name: "test2",
			args: args{
				image: [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			},
			want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
		{
			name: "test3",
			args: args{
				image: [][]int{{1}},
			},
			want: [][]int{{0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_832(tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Code_832() = %v, want %v", got, tt.want)
			}
		})
	}
}
