package leetcode

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 5, 0, 3, 5},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
