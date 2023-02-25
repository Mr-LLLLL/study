package leetcode

import "testing"

func TestCode_164(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{3, 6, 9, 1},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				nums: []int{10},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Code_164(tt.args.nums); got != tt.want {
				t.Errorf("Code_164() = %v, want %v", got, tt.want)
			}
		})
	}
}
