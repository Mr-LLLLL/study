package leetcode

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				n: 2,
			},
			want: []int{0, 1, 1},
		},
		{
			name: "test2",
			args: args{
				n: 5,
			},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
