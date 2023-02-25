package probles

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		left  int
		right int
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
				left:  0,
				right: 2,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				left:  2,
				right: 5,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				left:  0,
				right: 5,
			},
			want: -3,
		},
	}
	obj := NewNumArray([]int{-2, 0, 3, -5, 2, -1})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := obj.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
