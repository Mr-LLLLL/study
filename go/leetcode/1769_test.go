package leetcode

import (
	"reflect"
	"testing"
)

func Test_code_1769(t *testing.T) {
	type args struct {
		arr string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: "1001",
			},
			want: []int{3, 3, 3, 3},
		},
		{
			name: "test2",
			args: args{
				arr: "001011",
			},
			want: []int{11, 8, 5, 4, 3, 4},
		},
		{
			name: "test3",
			args: args{
				arr: "110110",
			},
			want: []int{8, 6, 6, 6, 8, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_1769(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("code1769() = %v, want %v", got, tt.want)
			}
		})
	}
}
