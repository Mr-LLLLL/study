package probles

import (
	"reflect"
	"sort"
	"testing"
)

func Test_smallestK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{1, 3, 5, 7, 2, 4, 6, 8},
				k:   4,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test2",
			args: args{
				arr: []int{3, 4, 8, 2, 9},
				k:   2,
			},
			want: []int{2, 3},
		},
		{
			name: "test3",
			args: args{
				arr: []int{3, 4, 8, 2, 9},
				k:   1,
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestK(tt.args.arr, tt.args.k)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestK() = %v, want %v", got, tt.want)
			}
		})
	}
}
