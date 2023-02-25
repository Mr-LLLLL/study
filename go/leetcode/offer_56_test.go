package leetcode

import (
	"reflect"
	"testing"
)

func TestCode_Offer_56(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{4, 1, 4, 6},
			},
			want: []int{1, 6},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 10, 4, 1, 4, 3, 3},
			},
			want: []int{2, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_Offer_56(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Code_Offer_56() = %v, want %v", got, tt.want)
			}
		})
	}
}
